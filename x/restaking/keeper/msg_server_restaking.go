package keeper

import (
	"context"
	"lightmos/x/restaking/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)
	logger.Info("carver|createValidator-start", "pubkey", validator.Pubkey.String())
	cosmosValidator := &stakingtypes.MsgCreateValidator{
		Description:       stakingtypes.Description(validator.Description),
		Commission:        stakingtypes.CommissionRates(validator.Commission),
		MinSelfDelegation: validator.MinSelfDelegation,
		DelegatorAddress:  validator.DelegatorAddress,
		ValidatorAddress:  validator.ValidatorAddress,
		Pubkey:            validator.Pubkey,
		Value:             cosmostypes.Coin(validator.Value),
	}

	// TODO create ibc packet
	res, err := k.stakingKeeper.RestakeValidator(goCtx, cosmosValidator)
	if err != nil {
		logger.Info("carver|createValidator-end", "err", err.Error())
	}
	return (*types.MsgCreateValidatorResponse)(res), err
}
