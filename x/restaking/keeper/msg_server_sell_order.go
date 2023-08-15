package keeper

import (
	"context"
	"errors"

	"lightmos/x/restaking/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendSellOrder(goCtx context.Context, msg *types.MsgSendSellOrder) (*types.MsgSendSellOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// If an order book doesn't exist, create
	pairIndex := types.OrderBookIndex(msg.Port, msg.ChannelID, msg.AmountDenom, msg.PriceDenom)
	sellOrderBook, found := k.GetSellOrderBook(ctx, pairIndex)
	if !found {
		sellOrderBook = types.NewSellOrderBook(msg.AmountDenom, msg.PriceDenom)
		sellOrderBook.Index = pairIndex
		k.SetSellOrderBook(ctx, sellOrderBook)
	}

	// The denom sending the sales order must be consistent with the amountDenom in the pair
	if sellOrderBook.AmountDenom != msg.AmountDenom ||
		k.stakingKeeper.BondDenom(ctx) != sellOrderBook.AmountDenom {
		return &types.MsgSendSellOrderResponse{}, errors.New("invalid amount denom")
	}

	// creator must be gov module account
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return &types.MsgSendSellOrderResponse{}, err
	}

	govAcc := k.accountKeeper.GetModuleAccount(ctx, "gov")
	if !sender.Equals(govAcc.GetAddress()) {
		return &types.MsgSendSellOrderResponse{}, errors.New("creator must be gov account")
	}

	var shouldMint sdkmath.Int
	if len(sellOrderBook.Book.Orders) == 0 {
		shouldMint = sdkmath.NewInt(int64(msg.Amount))
	} else {
		remainingAmt := sellOrderBook.Book.Orders[0].Amount
		if remainingAmt < msg.Amount {
			shouldMint = sdkmath.NewInt(int64(msg.Amount - remainingAmt))
		} else {
			shouldMint = sdkmath.NewInt(int64(msg.Amount))
		}
	}

	err = k.MintTokens(ctx, govAcc.GetAddress(), sdk.NewCoin(msg.AmountDenom, shouldMint))
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
		_, err := sellOrderBook.UpdateOrInsertOrder(msg.Creator, msg.Amount, msg.Price)
		if err != nil {
			return &types.MsgSendSellOrderResponse{}, err
		}

		// Save the new order book
		k.SetSellOrderBook(ctx, sellOrderBook)
	}

	return &types.MsgSendSellOrderResponse{}, nil
}
