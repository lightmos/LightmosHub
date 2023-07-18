package keeper

import (
	"lightmos/x/restaking/types"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

func (k Keeper) MintTokenForValidator(ctx sdk.Context, height int64) []abci.ValidatorUpdate {
	log := k.Logger(ctx)

	updateVal := make(map[string]sdk.Int)
	currVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height)
	log.Info("azh|restaking BeginBlock", "currVals len", len(currVals.Valset))

	preVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height-1)
	for _, val := range currVals.Valset {
		var exist bool
		for _, pre := range preVals.Valset {
			if val.OperatorAddress == pre.OperatorAddress {
				exist = true
				break
			}
		}
		if !exist {
			updateVal[val.OperatorAddress] = val.Tokens
		}
	}

	for val, amount := range updateVal {
		coins := sdk.NewInt64Coin("token", amount.Int64())
		valAdr, _ := sdk.ValAddressFromBech32(val)
		log.Info("azh|restaking BeginBlock", "accAddr", sdk.AccAddress(valAdr))
		k.MintTokens(ctx, sdk.AccAddress(valAdr), coins)
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

	packetAck.Succeed = false

	// mint token
	restaker, err := sdk.AccAddressFromBech32(data.Restaker)
	if err != nil {
		return packetAck, err
	}
	logger.Info("carver|recv restake packet", "restaker", restaker, "denom", sdk.DefaultBondDenom)

	err = k.MintTokens(ctx, restaker, sdk.NewCoin(sdk.DefaultBondDenom, data.Value.Amount))
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
		Pubkey:            data.Pubkey,
		Value:             sdk.Coin(data.Value),
	}

	// ## simple test restakeValidator ##
	_, err = k.stakingKeeper.RestakeValidator(ctx, cosmosValidator)
	if err != nil {
		logger.Error("carver|restake Validator err", err.Error())
		// if restake fail, burn tokens
		k.BurnTokens(ctx, restaker, sdk.NewCoin(data.Value.Denom, data.Value.Amount))
		return packetAck, err
	}

	packetAck.Succeed = true
	return packetAck, nil
}
