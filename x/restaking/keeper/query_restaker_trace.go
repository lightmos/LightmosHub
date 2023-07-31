package keeper

import (
	"context"

	"lightmos/x/restaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RestakerTrace(goCtx context.Context, req *types.QueryRestakerTraceRequest) (*types.QueryRestakerTraceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	rts := k.GetAllRestakerTrace(
		ctx,
	)
	var res []types.RestakerTrace
	for _, v := range rts {
		if v.Addr == req.AccAddr {
			res = append(res, v)
		}
	}

	return &types.QueryRestakerTraceResponse{RestakerTrace: res}, nil
}
