package keeper

import (
	"context"
	"errors"

	"lightmos/x/restaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendCreatePair(goCtx context.Context, msg *types.MsgSendCreatePair) (*types.MsgSendCreatePairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get an order book index
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.SourceDenom, msg.TargetDenom)
	demo := k.stakingKeeper.BondDenom(ctx)
	if demo == msg.TargetDenom {
		// If an order book is found, return an error
		_, found := k.GetBuyOrderBook(ctx, pairIndex)
		if found {
			return &types.MsgSendCreatePairResponse{}, errors.New("the pair already exist")
		}
		book := types.NewBuyOrderBook(msg.SourceDenom, msg.TargetDenom)
		book.Index = pairIndex
		k.SetBuyOrderBook(ctx, book)
	} else {
		// If an order book is found, return an error
		_, found := k.GetSellOrderBook(ctx, pairIndex)
		if found {
			return &types.MsgSendCreatePairResponse{}, errors.New("the pair already exist")
		}
		book := types.NewSellOrderBook(msg.SourceDenom, msg.TargetDenom)
		book.Index = pairIndex
		k.SetSellOrderBook(ctx, book)
	}
	return &types.MsgSendCreatePairResponse{}, nil
}
