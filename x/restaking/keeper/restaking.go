package keeper

import (
	"errors"
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

func (k Keeper) MintTokenForValidator(ctx sdk.Context, height int64) []abci.ValidatorUpdate {
	updateVal := make(map[string]sdk.Int)
	delVal := make(map[string]sdk.Int)
	currVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height)

	preVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height-1)
	for _, val := range currVals.Valset {
		var exist bool
		for _, pre := range preVals.Valset {
			if val.OperatorAddress == pre.OperatorAddress {
				if val.Tokens.LT(pre.Tokens) {
					valAddr, _ := sdk.ValAddressFromBech32(val.OperatorAddress)
					delVal[sdk.AccAddress(valAddr).String()] = pre.Tokens.Sub(val.Tokens)
				}
				exist = true
				break
			}
		}
		if !exist {
			valAddr, _ := sdk.ValAddressFromBech32(val.OperatorAddress)
			updateVal[sdk.AccAddress(valAddr).String()] = val.Tokens
		}
	}

	for _, pres := range preVals.Valset {
		var exist bool
		for _, vals := range currVals.Valset {
			if vals.OperatorAddress == pres.OperatorAddress {
				exist = true
				break
			}
		}
		if !exist {
			valAddr, _ := sdk.ValAddressFromBech32(pres.OperatorAddress)
			delVal[sdk.AccAddress(valAddr).String()] = pres.Tokens
		}
	}

	for val, amount := range updateVal {
		if _, found := k.GetValidatorToken(ctx, val); found {
			continue
		}
		coins := sdk.NewInt64Coin("token", amount.Int64())
		accAddr, _ := sdk.AccAddressFromBech32(val)
		k.MintTokens(ctx, accAddr, coins)
	}

	for val, amount := range delVal {
		if valToken, found := k.GetValidatorToken(ctx, val); found {
			amt := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), amount)
			accAdr, _ := sdk.AccAddressFromBech32(val)
			k.bankKeeper.SendCoinsFromAccountToModule(ctx, accAdr, "restaking", sdk.NewCoins(amt))
			valToken.Total -= amount.Uint64()
			valToken.Available = amount.Uint64()
			k.SetValidatorToken(ctx, valToken)
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

	vt := types.ValidatorToken{
		Address: data.Restaker,
		Total:   data.Value.Amount.Uint64(),
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
