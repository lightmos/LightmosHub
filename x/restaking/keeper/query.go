package keeper

import (
	"lightmos/x/restaking/types"
)

var _ types.QueryServer = Keeper{}
