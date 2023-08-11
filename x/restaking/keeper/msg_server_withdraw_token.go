package keeper

import (
	"context"

	"lightmos/x/restaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) WithdrawToken(goCtx context.Context, msg *types.MsgWithdrawToken) (*types.MsgWithdrawTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	accAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgWithdrawTokenResponse{}, err
	}

	valAddr := sdk.ValAddress(accAddr)
	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return &types.MsgWithdrawTokenResponse{}, types.ErrValidatorNotFound
	}

	if msg.Value.Amount.GT(validator.BondedTokens()) {
		return &types.MsgWithdrawTokenResponse{}, types.ErrExceedBondedAmount
	}

	balance, found := k.GetRemainingWithdrawToken(ctx, msg.Creator)
	if found && balance.LT(msg.Value.Amount) {
		// if recorded, should not be less than msg.value amount
		return &types.MsgWithdrawTokenResponse{}, types.ErrExceedBalance
	}

	err = k.MintTokens(ctx, accAddr, sdk.NewCoin(msg.Value.Denom, msg.Value.Amount))
	if err != nil {
		return &types.MsgWithdrawTokenResponse{}, err
	}

	if found {
		k.SetRemainingWithdrawToken(ctx, msg.Creator, balance.Sub(msg.Value.Amount))
	} else {
		k.SetRemainingWithdrawToken(ctx, msg.Creator, validator.BondedTokens().Sub(msg.Value.Amount))
	}

	return &types.MsgWithdrawTokenResponse{}, nil
}
