package keeper

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	restakingtypes "lightmos/x/restaking/types"
)

// TransmitUndelegatePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitUndelegatePacket(
	ctx sdk.Context,
	packetData restakingtypes.UndelegatePacketData,
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

// OnRecvUndelegatePacket processes packet reception
func (k Keeper) OnRecvUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData) (packetAck restakingtypes.UndelegatePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	recipientAcc := k.accountKeeper.GetModuleAccount(ctx, restakingtypes.ModuleName)
	if recipientAcc == nil {
		return packetAck, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", restakingtypes.ModuleName)
	}
	//destDenomFromVocher, flg := k.OriginalDenom(ctx, packet.DestinationPort, packet.DestinationChannel, data.Amount.Denom)
	//if !flg {
	//	return packetAck, errors.New("invalid denom")
	//}
	demo := k.stakingKeeper.BondDenom(ctx)
	amount := k.bankKeeper.GetBalance(ctx, recipientAcc.GetAddress(), demo)
	if err != nil {
		return packetAck, err
	}
	del, retire := k.DescHistory(ctx, "token", demo, data.ValidatorAddress, int32(data.Amount.Amount.Int64()))
	if !del || retire > int32(amount.Amount.Int64()) {
		return packetAck, errors.New("not exist buy sell")
	}
	coins := sdk.NewCoin(demo, data.Amount.Amount)
	if err = k.BurnTokens(ctx, recipientAcc.GetAddress(), coins); err != nil {
		return packetAck, err
	}

	k.Logger(ctx).Info("azh|OnRecvUndelegatePacket burn success")
	packetAck.Step = 1

	return packetAck, nil
}

// OnAcknowledgementUndelegatePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck restakingtypes.UndelegatePacketAck

		if err := restakingtypes.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		accAddr, err := sdk.AccAddressFromBech32(data.ValidatorAddress)
		if err != nil {
			return err
		}

		coins := sdk.NewCoin("token", data.Amount.Amount)

		return k.UnlockTokens(ctx, packet.SourcePort, packet.SourceChannel, accAddr, coins)

	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutUndelegatePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData) error {

	// TODO: packet timeout logic

	return nil
}
