package keeper

import (
	"errors"
	"lightmos/x/restaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
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
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}
	// TODO: packet reception logic
	accAddr, _ := sdk.AccAddressFromBech32(data.ValidatorAddress)
	vt, found := k.GetValidatorToken(ctx, accAddr.String())
	if !found {
		return packetAck, errors.New("not found")
	}
	if vt.Available.IsLT(*data.Amount) {
		return packetAck, errors.New("retire token is too large")
	} else {
		if !vt.Retire.IsZero() {
			coin := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), vt.Retire.Amount)
			if err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(coin)); err != nil {
				return packetAck, err
			}
			vt.Retire.Amount = sdk.ZeroInt()
		}

		if vt.Available.IsEqual(*data.Amount) && vt.Total.IsZero() {
			k.RemoveValidatorToken(ctx, accAddr.String())
		} else {
			avaCoin := vt.Available.Sub(*data.Amount)
			vt.Available = &avaCoin
			k.SetValidatorToken(ctx, vt)
		}
	}

	k.Logger(ctx).Info("azh|OnRecvUndelegatePacket burn success")
	packetAck.Step = 1
	return packetAck, nil
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

// OnTimeoutRetireSharePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutRetireSharePacket(ctx sdk.Context, packet channeltypes.Packet, data types.RetireSharePacketData) error {

	// TODO: packet timeout logic

	return nil
}
