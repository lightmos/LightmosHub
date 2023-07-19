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

	// var pkAny *codectypes.Any
	// if msg.Pubkey != nil {
	// 	var err error
	// 	if pkAny, err = codectypes.NewAnyWithValue(msg.Pubkey); err != nil {
	// 		return nil, err
	// 	}
	// }
	// Construct the packet
	packet, _ := types.NewRestakePacketData(
		msg.Creator, msg.DelegatorAddress, msg.ValidatorAddress,
		msg.Pubkey, msg.Value, msg.Description,
		msg.Commission, msg.MinSelfDelegation,
	)

	// packet.Description = msg.Description
	// packet.Commission = msg.Commission
	// packet.MinSelfDelegation = msg.MinSelfDelegation
	// packet.DelegatorAddress = msg.DelegatorAddress
	// packet.ValidatorAddress = msg.ValidatorAddress
	// packet.Pubkey = msg.Pubkey
	// packet.Value = msg.Value
	// packet.Restaker = msg.Creator

	cachedvalue := packet.Pubkey.GetCachedValue()
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)

	// # test # //cachedvalue
	k.Keeper.Logger(ctx).Info("carver|CreateValidator", "addr", creator,
		"denom", msg.Value.Denom, "amount", msg.Value.Amount, "cachedvalue", cachedvalue)

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
