package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"lightmos/x/restaking/types"
)

func (k msgServer) SendRetireShare(goCtx context.Context, msg *types.MsgSendRetireShare) (*types.MsgSendRetireShareResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.RetireSharePacketData

	packet.ValidatorAddress = msg.ValidatorAddress
	packet.Amount = msg.Amount
	// Transmit the packet
	_, err := k.TransmitRetireSharePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendRetireShareResponse{}, nil
}
