package types

import (
	types "lightmos/types"

	"cosmossdk.io/math"
)

// ValidateBasic is used for validating the packet
func (p RestakePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

func NewRestakePacketData(
	restaker, delegatorAddr, validatorAddr string,
	pubKey string, selfDelegation types.Coin, description Description,
	commission CommissionRates, minSelfDelegation math.Int,
	destinationChainId string,
) (RestakePacketData, error) {
	return RestakePacketData{
		Description:        description,
		Commission:         commission,
		MinSelfDelegation:  minSelfDelegation,
		DelegatorAddress:   delegatorAddr,
		ValidatorAddress:   validatorAddr,
		Pubkey:             pubKey,
		Value:              selfDelegation,
		Restaker:           restaker,
		DestinationChainId: destinationChainId,
	}, nil

}

// GetBytes is a helper for serialising
func (p RestakePacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_RestakePacket{&p}

	return modulePacket.Marshal()
}
