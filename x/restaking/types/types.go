package types

import (
	"errors"
	"sort"
)

const (
	MaxAmount = int32(10000000)
	MaxPrice  = int32(100000)
)

type Ordering int

const (
	Increasing Ordering = iota
	Decreasing
)

var (
	ErrMaxAmount     = errors.New("max amount reached")
	ErrMaxPrice      = errors.New("max price reached")
	ErrZeroAmount    = errors.New("amount is zero")
	ErrZeroPrice     = errors.New("price is zero")
	ErrOrderNotFound = errors.New("order not found")
)

func NewOrderBook() OrderBook {
	return OrderBook{
		IdCount: 0,
	}
}

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	book := NewOrderBook()
	return BuyOrderBook{
		AmountDenom: AmountDenom,
		PriceDenom:  PriceDenom,
		Book:        &book,
	}
}

func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrderBook {
	book := NewOrderBook()
	return SellOrderBook{
		AmountDenom: AmountDenom,
		PriceDenom:  PriceDenom,
		Book:        &book,
	}
}

func (book *OrderBook) appendOrder(creator string, amount int32, price int32, ordering Ordering) (int32, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return 0, err
	}

	// Initialize the order
	var order Order
	order.Id = book.GetNextOrderID()
	order.Creator = creator
	order.Amount = amount
	order.Price = price

	// Increment ID tracker
	book.IncrementNextOrderID()

	// Insert the order
	book.insertOrder(order, ordering)
	return order.Id, nil
}

func (book *OrderBook) updateOrAppendOrder(creator string, amount int32, price int32) (int32, error) {
	if err := checkAmountAndPrice(amount, price); err != nil {
		return 0, err
	}

	// Initialize the order
	var order Order

	order.Id = book.GetNextOrderID()
	order.Creator = creator
	order.Amount = amount
	order.Price = price

	foundCreator := func(order *Order) bool {
		for _, v := range book.Orders {
			if v.Creator == creator {
				return true
			}
			continue
		}
		return false
	}

	if len(book.Orders) > 0 && foundCreator(&order) {
		book.Orders[0] = &order
	} else {
		book.Orders = append(book.Orders, &order)
		// Increment ID tracker
		book.IncrementNextOrderID()
	}

	return order.Id, nil
}

func checkAmountAndPrice(amount int32, price int32) error {
	if amount == int32(0) {
		return ErrZeroAmount
	}
	if amount > MaxAmount {
		return ErrMaxAmount
	}

	if price == int32(0) {
		return ErrZeroPrice
	}
	if price > MaxPrice {
		return ErrMaxPrice
	}

	return nil
}

func (book OrderBook) GetNextOrderID() int32 {
	return book.IdCount
}

func (book *OrderBook) IncrementNextOrderID() {
	// Even numbers to have different ID than buy orders
	book.IdCount++
}

func (book *OrderBook) insertOrder(order Order, ordering Ordering) {
	if len(book.Orders) > 0 {
		var i int

		// get the index of the new order depending on the provided ordering
		if ordering == Increasing {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price > order.Price })
		} else {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price < order.Price })
		}

		// insert order
		orders := append(book.Orders, &order)
		copy(orders[i+1:], orders[i:])
		orders[i] = &order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, &order)
	}
}

func (b *BuyOrderBook) FillSellOrder(order Order) (
	remainingSellOrder Order,
	liquidated []Order,
	gain int32,
	filled bool,
) {
	var liquidatedList []Order
	totalGain := int32(0)
	remainingSellOrder = order

	// Liquidate as long as there is match
	for {
		var match bool
		var liquidation Order
		remainingSellOrder, liquidation, gain, match, filled = b.LiquidateFromSellOrder(
			remainingSellOrder,
		)
		if !match {
			break
		}

		// Update gains
		totalGain += gain

		// Update liquidated
		liquidatedList = append(liquidatedList, liquidation)

		if filled {
			break
		}
	}

	return remainingSellOrder, liquidatedList, totalGain, filled
}

func (b *BuyOrderBook) LiquidateFromSellOrder(order Order) (
	remainingSellOrder Order,
	liquidatedBuyOrder Order,
	gain int32,
	match bool,
	filled bool,
) {
	remainingSellOrder = order

	// No match if no order
	orderCount := len(b.Book.Orders)
	if orderCount == 0 {
		return order, liquidatedBuyOrder, gain, false, false
	}

	// Check if match
	highestBid := b.Book.Orders[orderCount-1]
	if order.Price > highestBid.Price {
		return order, liquidatedBuyOrder, gain, false, false
	}

	liquidatedBuyOrder = *highestBid

	// Check if sell order can be entirely filled
	if highestBid.Amount >= order.Amount {
		remainingSellOrder.Amount = 0
		liquidatedBuyOrder.Amount = order.Amount
		gain = order.Amount * highestBid.Price

		// Remove the highest bid if it has been entirely liquidated
		highestBid.Amount -= order.Amount
		if highestBid.Amount == 0 {
			b.Book.Orders = b.Book.Orders[:orderCount-1]
		} else {
			b.Book.Orders[orderCount-1] = highestBid
		}

		return remainingSellOrder, liquidatedBuyOrder, gain, true, true
	}

	// Not entirely filled
	gain = highestBid.Amount * highestBid.Price
	b.Book.Orders = b.Book.Orders[:orderCount-1]
	remainingSellOrder.Amount -= highestBid.Amount

	return remainingSellOrder, liquidatedBuyOrder, gain, true, false
}

func (s *SellOrderBook) AppendOrder(creator string, amount int32, price int32) (int32, error) {
	return s.Book.appendOrder(creator, amount, price, Decreasing)
}

func (s *SellOrderBook) UpdateOrInsertOrder(creator string, amount int32, price int32) (int32, error) {
	return s.Book.updateOrAppendOrder(creator, amount, price)
}

func (s *SellOrderBook) FillBuyOrder(order Order) (liquidated Order, match bool) {
	// Liquidate as long as there is match
	orderCount := len(s.Book.Orders)
	if orderCount == 0 {
		return
	}
	for i := 0; i < orderCount; i++ {
		currentAsk := s.Book.Orders[i]
		if order.Price == currentAsk.Price {
			if currentAsk.Amount >= order.Amount {
				match = true
				currentAsk.Amount -= order.Amount
				liquidated.Creator = currentAsk.Creator
				liquidated.Amount = order.Amount
				liquidated.Price = order.Price
				if currentAsk.Amount == 0 {
					s.Book.Orders = append(s.Book.Orders[:i], s.Book.Orders[i+1:]...)
				} else {
					s.Book.Orders[i] = currentAsk
				}
				break
			}
		}
	}
	return
}

func (s *SellOrderBook) LiquidateFromBuyOrder(order Order) (
	remainingBuyOrder Order,
	liquidatedSellOrder Order,
	purchase int32,
	match bool,
	filled bool,
) {
	remainingBuyOrder = order

	// No match if no order
	orderCount := len(s.Book.Orders)
	if orderCount == 0 {
		return order, liquidatedSellOrder, purchase, false, false
	}

	// Check if match
	lowestAsk := s.Book.Orders[orderCount-1]
	if order.Price < lowestAsk.Price {
		return order, liquidatedSellOrder, purchase, false, false
	}

	liquidatedSellOrder = *lowestAsk

	// Check if buy order can be entirely filled
	if lowestAsk.Amount >= order.Amount {
		remainingBuyOrder.Amount = 0
		liquidatedSellOrder.Amount = order.Amount
		purchase = order.Amount

		// Remove lowest ask if it has been entirely liquidated
		lowestAsk.Amount -= order.Amount
		if lowestAsk.Amount == 0 {
			s.Book.Orders = s.Book.Orders[:orderCount-1]
		} else {
			s.Book.Orders[orderCount-1] = lowestAsk
		}

		return remainingBuyOrder, liquidatedSellOrder, purchase, true, true
	}

	// Not entirely filled
	purchase = lowestAsk.Amount
	s.Book.Orders = s.Book.Orders[:orderCount-1]
	remainingBuyOrder.Amount -= lowestAsk.Amount

	return remainingBuyOrder, liquidatedSellOrder, purchase, true, false
}

func (b *BuyOrderBook) AppendOrder(creator string, amount int32, price int32) (int32, error) {
	return b.Book.appendOrder(creator, amount, price, Increasing)
}

func (book OrderBook) GetOrderFromID(id int32) (Order, error) {
	for _, order := range book.Orders {
		if order.Id == id {
			return *order, nil
		}
	}

	return Order{}, ErrOrderNotFound
}

func (book *OrderBook) RemoveOrderFromID(id int32) error {
	for i, order := range book.Orders {
		if order.Id == id {
			book.Orders = append(book.Orders[:i], book.Orders[i+1:]...)
			return nil
		}
	}

	return ErrOrderNotFound
}
