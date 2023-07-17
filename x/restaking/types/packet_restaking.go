package types

// ValidateBasic is used for validating the packet
func (p RestakePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p RestakePacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_RestakePacket{&p}

	return modulePacket.Marshal()
}
