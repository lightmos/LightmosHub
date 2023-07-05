package types

import fmt "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "restaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_restaking"

	// Version defines the current version the IBC module supports
	Version = "restaking-1"

	// PortID is the default port id that module binds to
	PortID = "restaking"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("restaking-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func OrderBookIndex(portID string, channelID string, sourceDenom string, targetDenom string) string {
	return fmt.Sprintf("%s-%s-%s-%s", portID, channelID, sourceDenom, targetDenom)
}
