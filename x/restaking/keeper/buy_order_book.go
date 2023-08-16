package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lightmos/x/restaking/types"
)

// SetBuyOrderBook set a specific buyOrderBook in the store from its index
func (k Keeper) SetBuyOrderBook(ctx sdk.Context, buyOrderBook types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	b := k.cdc.MustMarshal(&buyOrderBook)
	store.Set(types.BuyOrderBookKey(
		buyOrderBook.Index,
	), b)
}

// GetBuyOrderBook returns a buyOrderBook from its index
func (k Keeper) GetBuyOrderBook(
	ctx sdk.Context,
	index string,

) (val types.BuyOrderBook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))

	b := store.Get(types.BuyOrderBookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBuyOrderBook removes a buyOrderBook from the store
func (k Keeper) RemoveBuyOrderBook(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	store.Delete(types.BuyOrderBookKey(
		index,
	))
}

// GetAllBuyOrderBook returns all buyOrderBook
func (k Keeper) GetAllBuyOrderBook(ctx sdk.Context) (list []types.BuyOrderBook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BuyOrderBookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BuyOrderBook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) storeHistory(ctx sdk.Context, orderHistory types.DoneChanHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderDoneHistoryPrefix))
	b := k.cdc.MustMarshal(&orderHistory)
	store.Set(types.DemoDoneHistory(
		orderHistory.GetSrcDemo()+orderHistory.GetDstDemo(),
	), b)
}

func (k Keeper) RemoveHistory(ctx sdk.Context, srcDemo, destDemo string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderDoneHistoryPrefix))
	store.Delete(types.DemoDoneHistory(srcDemo + destDemo))
}

func (k Keeper) GetDemoHistory(
	ctx sdk.Context,
	srcDemo string,
	destDemo string,
) (val types.DoneChanHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrderDoneHistoryPrefix))

	b := store.Get(types.DemoDoneHistory(
		srcDemo + destDemo,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateDoneHistory(ctx sdk.Context, srcDemo, destDemo, buyer, seller string, price int32, amount int32) error {
	orderDoneList := &types.OrderDoneList{
		Amount: amount,
		Price:  price,
	}
	demoList := types.OrderDemoList{
		Buyer:  buyer,
		Seller: seller,
	}
	history, ok := k.GetDemoHistory(ctx, srcDemo, destDemo)
	if ok {
		for _, doneList := range history.OrderDemo {
			if doneList.Buyer == buyer && doneList.Seller == seller {
				var orderFlag bool
				for _, doneOrder := range doneList.OrderDoneList {
					if doneOrder.Price == price {
						doneOrder.Amount += amount
						orderFlag = true
						break
					}
				}
				if !orderFlag {
					doneList.OrderDoneList = append(doneList.OrderDoneList, orderDoneList)
				}

				k.storeHistory(ctx, history)
				return nil
			}
		}

		demoList.OrderDoneList = append(demoList.OrderDoneList, orderDoneList)
		history.OrderDemo = append(history.OrderDemo, &demoList)
		k.storeHistory(ctx, history)
		return nil
	}

	history.SrcDemo = srcDemo
	history.DstDemo = destDemo
	demoList.OrderDoneList = append(demoList.OrderDoneList, orderDoneList)
	history.OrderDemo = append(history.OrderDemo, &demoList)
	k.storeHistory(ctx, history)
	return nil
}

func (k Keeper) DescHistory(ctx sdk.Context, srcDemo, destDemo, creator string, amount int32) (found bool, retire int32) {
	history, ok := k.GetDemoHistory(ctx, srcDemo, destDemo)
	if ok {
		var listFlag bool
	loop:
		for _, doneList := range history.OrderDemo {
			if doneList.Buyer == creator {
				listFlag = true
				for i := 0; i < len(doneList.OrderDoneList); i++ {
					currentOrder := doneList.OrderDoneList[i]
					if currentOrder.Amount >= amount {
						currentOrder.Amount -= amount
						if currentOrder.Amount == amount {
							doneList.OrderDoneList = append(doneList.OrderDoneList[:i], doneList.OrderDoneList[i+1:]...)
						} else {
							doneList.OrderDoneList[i] = currentOrder
						}
						retire += amount * currentOrder.Price
						break loop
					} else {
						doneList.OrderDoneList = append(doneList.OrderDoneList[:i], doneList.OrderDoneList[i+1:]...)
						retire += currentOrder.Amount * currentOrder.Price
					}
				}
			}
		}

		if listFlag {
			k.storeHistory(ctx, history)
			return true, retire
		}
	}
	return false, 0
}
