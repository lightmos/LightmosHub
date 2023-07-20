package keeper

import (
	"lightmos/x/restaking/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRestakeValidatorTrace set a specific validatorTrace in the store from its index
func (k Keeper) SetRestakeValidatorTrace(ctx sdk.Context, restaker string, destinationChainId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))
	store.Set(types.RestakeServiceKey(restaker), []byte(destinationChainId))
}

func (k Keeper) GetRestakeValidatorTrace(
	ctx sdk.Context,
	restaker string,
) (destinationChainId string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))

	b := store.Get(types.RestakeServiceKey(
		restaker,
	))
	if b == nil {
		return destinationChainId, false
	}

	return string(b), true
}

// RemoveRestakeValidatorTrace will set restake validator trace as empty
func (k Keeper) RemoveRestakeValidatorTrace(
	ctx sdk.Context,
	restaker string,
) bool {
	if _, found := k.GetRestakeValidatorTrace(ctx, restaker); found {
		k.SetRestakeValidatorTrace(ctx, restaker, "")
		return true
	}
	return false
}
