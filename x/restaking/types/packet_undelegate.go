package types

// ValidateBasic is used for validating the packet
func (p UndelegatePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p UndelegatePacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_UndelegatePacket{&p}

	return modulePacket.Marshal()
}
