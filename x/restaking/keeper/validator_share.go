package keeper

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintTokenForValidator(ctx sdk.Context, height int64) []abci.ValidatorUpdate {
	log := k.Logger(ctx)

	updateVal := make(map[string]sdk.Int)
	currVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height)
	log.Info("azh|restaking BeginBlock", "currVals len", len(currVals.Valset))

	preVals, _ := k.stakingKeeper.GetHistoricalInfo(ctx, height-1)
	for _, val := range currVals.Valset {
		var exist bool
		for _, pre := range preVals.Valset {
			if val.OperatorAddress == pre.OperatorAddress {
				exist = true
				break
			}
		}
		if !exist {
			updateVal[val.OperatorAddress] = val.Tokens
		}
	}

	for val, amount := range updateVal {
		coins := sdk.NewInt64Coin("token", amount.Int64())
		valAdr, _ := sdk.ValAddressFromBech32(val)
		log.Info("azh|restaking BeginBlock", "accAddr", sdk.AccAddress(valAdr))
		k.MintTokens(ctx, sdk.AccAddress(valAdr), coins)
	}

	return []abci.ValidatorUpdate{}
}
