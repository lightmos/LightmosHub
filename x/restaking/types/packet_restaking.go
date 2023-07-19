package types

import (
	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

var (
	_ codectypes.UnpackInterfacesMessage = (*RestakePacketData)(nil)
)

// ValidateBasic is used for validating the packet
func (p RestakePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

func NewRestakePacketData(
	restaker, delegatorAddr, validatorAddr string,
	pubKey *types.Any, selfDelegation Coin, description Description,
	commission CommissionRates, minSelfDelegation math.Int,
) (RestakePacketData, error) {
	var pkAny *codectypes.Any
	if pubKey != nil {
		var err error
		if pkAny, err = codectypes.NewAnyWithValue(pubKey); err != nil {
			return RestakePacketData{}, err
		}
	}
	return RestakePacketData{
		Description:       description,
		Commission:        commission,
		MinSelfDelegation: minSelfDelegation,
		DelegatorAddress:  delegatorAddr,
		ValidatorAddress:  validatorAddr,
		Pubkey:            pkAny,
		Value:             selfDelegation,
		Restaker:          restaker,
	}, nil

}

// GetBytes is a helper for serialising
func (p RestakePacketData) GetBytes() ([]byte, error) {
	var modulePacket RestakingPacketData

	modulePacket.Packet = &RestakingPacketData_RestakePacket{&p}

	return modulePacket.Marshal()
}

func (rpd RestakePacketData) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var pubKey cryptotypes.PubKey
	return unpacker.UnpackAny(rpd.Pubkey, &pubKey)
}
