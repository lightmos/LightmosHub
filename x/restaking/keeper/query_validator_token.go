package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"lightmos/x/restaking/types"
)

func (k Keeper) ValidatorTokenAll(goCtx context.Context, req *types.QueryAllValidatorTokenRequest) (*types.QueryAllValidatorTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var validatorTokens []types.ValidatorToken
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	validatorTokenStore := prefix.NewStore(store, types.KeyPrefix(types.ValidatorTokenKey))

	pageRes, err := query.Paginate(validatorTokenStore, req.Pagination, func(key []byte, value []byte) error {
		var validatorToken types.ValidatorToken
		if err := k.cdc.Unmarshal(value, &validatorToken); err != nil {
			return err
		}

		validatorTokens = append(validatorTokens, validatorToken)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllValidatorTokenResponse{ValidatorToken: validatorTokens, Pagination: pageRes}, nil
}

func (k Keeper) ValidatorToken(goCtx context.Context, req *types.QueryGetValidatorTokenRequest) (*types.QueryGetValidatorTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	validatorToken, found := k.GetValidatorToken(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetValidatorTokenResponse{ValidatorToken: validatorToken}, nil
}
