package types

// ValidateBasic is used for validating the packet
func (p BuyOrderPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p BuyOrderPacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_BuyOrderPacket{&p}

	return modulePacket.Marshal()
}
