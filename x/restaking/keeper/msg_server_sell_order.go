package keeper

import (
	"context"
	"errors"

	"lightmos/x/restaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendSellOrder(goCtx context.Context, msg *types.MsgSendSellOrder) (*types.MsgSendSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// If an order book doesn't exist, throw an error
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	sellOrderBook, found := k.GetSellOrderBook(ctx, pairIndex)
	if !found {
		return &types.MsgSendSellOrderResponse{}, errors.New("the pair doesn't exist")
	}

	// The denom sending the sales order must be consistent with the amountDenom in the pair
	if sellOrderBook.AmountDenom != msg.AmountDenom ||
		sellOrderBook.AmountDenom != k.stakingKeeper.BondDenom(ctx) {
		return &types.MsgSendSellOrderResponse{}, errors.New("invalid amount denom")
	}

	// Get sender's address
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgSendSellOrderResponse{}, err
	}

	// Use SafeBurn to ensure no new native tokens are minted
	if err := k.SafeBurn(ctx, msg.Port, msg.ChannelID, sender, msg.AmountDenom, msg.Amount); err != nil {
		return &types.MsgSendSellOrderResponse{}, err
	}

	// Save the voucher received on the other chain, to have the ability to resolve it into the original denom
	k.SaveVoucherDenom(ctx, msg.Port, msg.ChannelID, msg.AmountDenom)

	// Append the remaining amount of the order
	if msg.Amount > 0 {
		_, err := sellOrderBook.AppendOrder(msg.Creator, msg.Amount, msg.Price)
		if err != nil {
			return &types.MsgSendSellOrderResponse{}, err
		}

		// Save the new order book
		k.SetSellOrderBook(ctx, sellOrderBook)
	}

	return &types.MsgSendSellOrderResponse{}, nil
}
