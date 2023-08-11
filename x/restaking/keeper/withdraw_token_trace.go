package keeper

import (
	"lightmos/x/restaking/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
)

// SetWithdrawTokenTrace set a specific withdrawTokenTrace in the store from its index
// TODO if update validators cant handle bonded diff when appen delegation, record validator, current unboned, remaining amount
func (k Keeper) SetRemainingWithdrawToken(ctx sdk.Context, validator string, amount sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawTokenKeyPrefix))
	store.Set(types.WithdrawTokenKey(validator), []byte(amount.String()))
}

func (k Keeper) GetRemainingWithdrawToken(
	ctx sdk.Context,
	validator string,
) (amount sdk.Int, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawTokenKeyPrefix))

	b := store.Get(types.WithdrawTokenKey(
		validator,
	))
	if b == nil {
		return sdk.ZeroInt(), false
	}

	amount, _ = sdkmath.NewIntFromString(string(b))
	return amount, true
}
