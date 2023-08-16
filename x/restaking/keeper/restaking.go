package keeper

import (
	"errors"
	ty "lightmos/types"
	"lightmos/x/restaking/types"

	abci "github.com/cometbft/cometbft/abci/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

func (k Keeper) RestakeUndelegate(ctx sdk.Context) []abci.ValidatorUpdate {
	shareVals := k.GetAllValidatorToken(ctx)
	for _, shareVal := range shareVals {
		accAddr, _ := sdk.AccAddressFromBech32(shareVal.Address)
		demo := k.stakingKeeper.BondDenom(ctx)
		if val, found := k.stakingKeeper.GetValidator(ctx, sdk.ValAddress(accAddr)); found {
			if val.Tokens.LT(shareVal.Total.Amount) {
				retire := shareVal.Total.Amount.Sub(val.Tokens)
				amt := sdk.NewCoin(demo, retire)
				bal := k.bankKeeper.GetBalance(ctx, accAddr, demo)
				if bal.IsGTE(amt) {
					k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(amt))
					del, retireToken := k.DescHistory(ctx, k.stakingKeeper.BondDenom(ctx), "token", shareVal.Address, int32(retire.Int64()))
					if !del {
						return []abci.ValidatorUpdate{}
					}
					totalCoin := ty.NewCoin(demo, val.Tokens)
					shareVal.Total = &totalCoin
					retireCoin := ty.NewCoin(demo, retire)
					shareVal.Retire = &retireCoin
					avaCoin := ty.NewCoin("token", sdk.NewInt(int64(retireToken)))
					shareVal.Available = &avaCoin
					k.SetValidatorToken(ctx, shareVal)
				}
			}
		} else if !shareVal.Total.IsZero() {
			amt := sdk.NewCoin(demo, shareVal.Total.Amount)
			bal := k.bankKeeper.GetBalance(ctx, accAddr, demo)
			if bal.IsGTE(amt) {
				k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAddr, types.ModuleName, sdk.NewCoins(amt))
				del, retireToken := k.DescHistory(ctx, k.stakingKeeper.BondDenom(ctx), "token", shareVal.Address, int32(shareVal.Total.Amount.Int64()))
				if !del {
					return []abci.ValidatorUpdate{}
				}
				retireCoin := ty.NewCoin(demo, shareVal.Retire.Amount.Add(shareVal.Total.Amount))
				shareVal.Retire = &retireCoin
				avaCoin := ty.NewCoin("token", shareVal.Available.Amount.Add(sdk.NewInt(int64(retireToken))))
				shareVal.Available = &avaCoin
				shareVal.Total = nil
				k.SetValidatorToken(ctx, shareVal)
			}
		}
	}
	return []abci.ValidatorUpdate{}
}

// TransmitRestakingPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitRestakingPacket(
	ctx sdk.Context,
	packetData types.RestakePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()

	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvRestakePacket processes packet reception
func (k Keeper) OnRecvRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) (packetAck types.RestakePacketDataAck, err error) {
	goctx := sdk.UnwrapSDKContext(ctx)
	logger := k.Logger(goctx)

	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	var pk cryptotypes.PubKey
	if err := k.cdc.UnmarshalInterfaceJSON([]byte(data.Pubkey), &pk); err != nil {
		return packetAck, err
	}

	var pkAny *codectypes.Any
	if pk != nil {
		var err error
		if pkAny, err = codectypes.NewAnyWithValue(pk); err != nil {
			return packetAck, err
		}
	}

	packetAck.Succeed = false

	// mint token
	destDenomFromVocher, flg := k.OriginalDenom(ctx, packet.DestinationPort, packet.DestinationChannel, data.Value.Denom)
	if !flg {
		return packetAck, errors.New("invalid denom")
	}
	restaker, err := sdk.AccAddressFromBech32(data.Restaker)
	if err != nil {
		return packetAck, err
	}
	logger.Info("carver|recv restake packet", "restaker", restaker, "denom", data.Value.Denom,
		"destDenomFromVocher", destDenomFromVocher, "data", data)

	err = k.MintTokens(ctx, restaker, sdk.NewCoin(destDenomFromVocher, data.Value.Amount))
	if err != nil {
		return packetAck, err
	}

	// restake validator
	cosmosValidator := &stakingtypes.MsgCreateValidator{
		Description:       stakingtypes.Description(data.Description),
		Commission:        stakingtypes.CommissionRates(data.Commission),
		MinSelfDelegation: data.MinSelfDelegation,
		DelegatorAddress:  data.DelegatorAddress,
		ValidatorAddress:  data.ValidatorAddress,
		Pubkey:            pkAny,
		Value:             sdk.NewCoin(destDenomFromVocher, data.Value.Amount),
	}

	// ## simple test restakeValidator ##
	_, err = k.stakingKeeper.RestakeValidator(ctx, cosmosValidator)
	if err != nil {
		logger.Error("carver|restake Validator err", "err", err.Error())
		// if restake fail, burn tokens
		k.BurnTokens(ctx, restaker, sdk.NewCoin(data.Value.Denom, data.Value.Amount))
		return packetAck, err
	}

	coin := ty.NewCoin(k.stakingKeeper.BondDenom(ctx), data.Value.Amount)
	vt := types.ValidatorToken{
		Address: data.Restaker,
		Total:   &coin,
	}
	k.SetValidatorToken(ctx, vt)
	logger.Info("carver|recv restake handle succeed", "restaker", restaker, "denom", data.Value.Denom)
	packetAck.Succeed = true
	return packetAck, nil
}

// OnAcknowledgementRestakePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		ctx.Logger().Info("caver|OnAcknowledgementRestakePacket err")
		return k.refundPacketToken(ctx, packet, data)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.RestakePacketDataAck
		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		ctx.Logger().Info("caver|OnAcknowledgementRestakePacket succeed")
		// save restake validator trace
		k.SetRestakeValidatorTrace(ctx, data.Restaker, data.DestinationChainId)
		return nil
	default:
		return nil
	}
}

func (k Keeper) refundPacketToken(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) error {
	// In case of error we unlock the native token
	receiver, err := sdk.AccAddressFromBech32(data.Restaker)
	if err != nil {
		return err
	}

	if err := k.UnlockTokens(
		ctx,
		packet.SourcePort,
		packet.SourceChannel,
		receiver,
		sdk.Coin(data.Value),
	); err != nil {
		return err
	}

	return nil
}

// OnTimeoutRestakePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutRestakePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RestakePacketData) error {
	return k.refundPacketToken(ctx, packet, data)
}
