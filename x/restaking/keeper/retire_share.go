package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"lightmos/x/restaking/types"
)

// TransmitRetireSharePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitRetireSharePacket(
	ctx sdk.Context,
	packetData types.RetireSharePacketData,
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

// OnRecvRetireSharePacket processes packet reception
func (k Keeper) OnRecvRetireSharePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RetireSharePacketData) (packetAck types.RetireSharePacketAck, err error) {
	// validate packet data upon receiving
	log := k.Logger(ctx)
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}
	// TODO: packet reception logic

	//accAddr, _ := sdk.AccAddressFromBech32(data.ValidatorAddress)
	accAddr, _ := sdk.AccAddressFromBech32("cosmos19se8lq7vs33hnvd6qg6wanad36r64r88uh56r6")
	valAdr := sdk.ValAddress(accAddr)
	log.Info("azh|OnRecvRetireSharePacket", "accAddr", accAddr)
	del, found := k.stakingKeeper.GetDelegation(ctx, accAddr, valAdr)
	if !found {
		return packetAck, types.ErrNoDelegation
	}
	log.Info("azh|OnRecvRetireSharePacket", "demo", k.stakingKeeper.BondDenom(ctx), "shares", del.Shares)
	if k.stakingKeeper.BondDenom(ctx) == "stake" {
		shares, err := k.stakingKeeper.ValidateUnbondAmount(
			ctx, accAddr, valAdr, data.Amount.Amount,
		)
		if err != nil {
			return packetAck, err
		}
		log.Info("azh|OnRecvRetireSharePacket", "undelegate", shares)
		_, err = k.stakingKeeper.Undelegate(ctx, accAddr, valAdr, shares)
		return packetAck, err
	}

	return packetAck, errors.New("does`t exit token or doesn`t have")
}

// OnAcknowledgementRetireSharePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementRetireSharePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RetireSharePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.RetireSharePacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutRetireSharePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutRetireSharePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RetireSharePacketData) error {

	// TODO: packet timeout logic

	return nil
}
