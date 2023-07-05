package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lightmos/testutil/keeper"
	"lightmos/x/restaking/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RestakingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
