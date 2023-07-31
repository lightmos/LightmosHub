package keeper

import (
	"lightmos/x/restaking/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRestakerTrace set a specific validatorTrace in the store from its index
func (k Keeper) SetRestakerTrace(ctx sdk.Context, rt types.RestakerTrace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))
	b := k.cdc.MustMarshal(&rt)
	store.Set(types.RestakerTraceKey(rt.Addr, rt.DestChainId), b)
}

// GetRestakerTrace returns a restakerTrace from its index
func (k Keeper) GetRestakerTrace(
	ctx sdk.Context,
	addr, destChainId string,
) (val types.RestakerTrace, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))

	b := store.Get(types.RestakerTraceKey(
		addr,
		destChainId,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRestakerTrace removes a restakerTrace from the store
func (k Keeper) RemoveRestakerTrace(
	ctx sdk.Context,
	addr, destChainId string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))
	store.Delete(types.RestakerTraceKey(
		addr, destChainId,
	))
}

// GetAllRestakerTrace returns all restakerTrace
func (k Keeper) GetAllRestakerTrace(ctx sdk.Context) (list []types.RestakerTrace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RestakeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RestakerTrace
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
