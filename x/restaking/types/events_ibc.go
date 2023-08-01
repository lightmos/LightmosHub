package types

// IBC events
const (
	EventTypeTimeout           = "timeout"
	EventTypeCreatePairPacket  = "createPair_packet"
	EventTypeSellOrderPacket   = "sellOrder_packet"
	EventTypeBuyOrderPacket    = "buyOrder_packet"
	EventTypeRestakePacket     = "restake_packet"
	EventTypeRetireSharePacket = "retireShare_packet"

	EventTypeUndelegatePacket = "undelegate_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
