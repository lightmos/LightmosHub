package keeper

import (
	"context"
	"lightmos/x/restaking/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Construct the packet
	packet, _ := types.NewRestakePacketData(
		msg.Creator, msg.DelegatorAddress, msg.ValidatorAddress,
		msg.Pubkey, msg.Value, msg.Description,
		msg.Commission, msg.MinSelfDelegation,
		msg.DestinationChainId,
	)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	k.Keeper.Logger(ctx).Info("carver|CreateValidator", "addr", creator,
		"denom", msg.Value.Denom, "amount", msg.Value.Amount)

	// Lock the tokens
	if err := k.LockTokens(ctx, msg.Port, msg.ChannelID, creator,
		sdk.NewCoin(msg.Value.Denom,
			sdkmath.NewInt(msg.Value.Amount.Int64()))); err != nil {
		return &types.MsgCreateValidatorResponse{}, err
	}

	// Transmit the packet
	_, err := k.TransmitRestakingPacket(
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

	return &types.MsgCreateValidatorResponse{}, err
}
