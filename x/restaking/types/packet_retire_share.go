package types

// ValidateBasic is used for validating the packet
func (p RetireSharePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p RetireSharePacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_RetireSharePacket{&p}

	return modulePacket.Marshal()
}
