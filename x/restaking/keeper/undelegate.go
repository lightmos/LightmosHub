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
	log := k.Logger(ctx)
	accAddr, _ := sdk.AccAddressFromBech32(data.ValidatorAddress)
	valAdr := sdk.ValAddress(accAddr)
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	if bondDenom != data.Amount.Denom {
		return packetAck, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", data.Amount.Denom, bondDenom)
	}

	shares, err := k.stakingKeeper.ValidateUnbondAmount(
		ctx, accAddr, valAdr, data.Amount.Amount,
	)
	if err != nil {
		return packetAck, err
	}
	endTime, err := k.stakingKeeper.Undelegate(ctx, accAddr, valAdr, shares)
	if err != nil {
		return packetAck, err
	}
	log.Info("azh|OnRecvRetireSharePacket", "undelegate", shares, "endTIme", endTime)
	packetAck.Step = 1

	return packetAck, nil
}

// OnAcknowledgementUndelegatePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData, ack channeltypes.Acknowledgement) error {
	log := k.Logger(ctx)
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		return errors.New(dispatchedAck.Error)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck restakingtypes.UndelegatePacketAck

		if err := restakingtypes.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		log.Info("azh|OnAcknowledgementRetireSharePacket", "dispatchedAck", packetAck.Step)
		if packetAck.Step == 1 {
			log.Info("azh|OnAcknowledgementRetireSharePacket unbound")
		}
		return nil
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
