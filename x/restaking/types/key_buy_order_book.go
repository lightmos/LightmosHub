package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BuyOrderBookKeyPrefix is the prefix to retrieve all BuyOrderBook
	BuyOrderBookKeyPrefix  = "BuyOrderBook/value/"
	OrderDoneHistoryPrefix = "OrderDoneHistory/value/"
)

// BuyOrderBookKey returns the store key to retrieve a BuyOrderBook from the index fields
func BuyOrderBookKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}

func DemoDoneHistory(
	demo string,
) []byte {
	var key []byte

	indexBytes := []byte(demo)
	key = append(key, indexBytes...)
	key = append(key, []byte("/done/")...)

	return key
}
