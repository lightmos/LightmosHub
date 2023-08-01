package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"lightmos/x/restaking/types"
)

func (k msgServer) ChangePairState(goCtx context.Context, msg *types.MsgChangePairState) (*types.MsgChangePairStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.SourceDenom, msg.TargetDenom)

	// If an order book is found, return an error
	if msg.Author == "buyer" {
		book, found := k.GetBuyOrderBook(ctx, pairIndex)
		if !found {
			return &types.MsgChangePairStateResponse{}, errors.New("the pair doesn`t exist")
		}

		if book.Book == nil {
			newBook := types.NewOrderBook()
			newBook.Orders = append(newBook.Orders, &types.Order{
				Id:      0,
				Creator: msg.Creator,
				Price:   msg.Price,
				Agree:   true,
			})
			book.Book = &newBook
		} else {
			for _, orders := range book.Book.Orders {
				if orders.Price == msg.Price {
					orders.Agree = true
				}
			}
		}
		// Save the new order book
		k.SetBuyOrderBook(ctx, book)
	} else if msg.Author == "seller" {
		book, found := k.GetSellOrderBook(ctx, pairIndex)
		if !found {
			return &types.MsgChangePairStateResponse{}, errors.New("the pair doesn`t exist")
		}

		if book.Book == nil {
			newBook := types.NewOrderBook()
			newBook.Orders = append(newBook.Orders, &types.Order{
				Id:      0,
				Creator: msg.Creator,
				Price:   msg.Price,
				Agree:   true,
			})
			book.Book = &newBook
		} else {
			for _, orders := range book.Book.Orders {
				if orders.Price == msg.Price {
					orders.Agree = true
				}
			}
		}
		// Save the new order book
		k.SetSellOrderBook(ctx, book)
	}

	return &types.MsgChangePairStateResponse{}, nil
}
