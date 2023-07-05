package keeper

import (
	"context"
	"lightmos/x/restaking/types"
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

func (k msgServer) CreateValidator(goCtx context.Context, validator *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	return nil, nil
}
