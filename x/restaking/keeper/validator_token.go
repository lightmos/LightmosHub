package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lightmos/x/restaking/types"
)

// SetValidatorToken set a specific validatorToken in the store
func (k Keeper) SetValidatorToken(ctx sdk.Context, validatorToken types.ValidatorToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorTokenKey))
	b := k.cdc.MustMarshal(&validatorToken)
	store.Set([]byte(validatorToken.Address), b)
}

// GetValidatorToken returns a validatorToken from its id
func (k Keeper) GetValidatorToken(ctx sdk.Context, address string) (val types.ValidatorToken, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorTokenKey))
	b := store.Get([]byte(address))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidatorToken removes a validatorToken from the store
func (k Keeper) RemoveValidatorToken(ctx sdk.Context, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorTokenKey))
	store.Delete([]byte(address))
}

// GetAllValidatorToken returns all validatorToken
func (k Keeper) GetAllValidatorToken(ctx sdk.Context) (list []types.ValidatorToken) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorTokenKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ValidatorToken
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetValidatorTokenIDBytes returns the byte representation of the ID
func GetValidatorTokenIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetValidatorTokenIDFromBytes returns ID in uint64 format from a byte array
func GetValidatorTokenIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
