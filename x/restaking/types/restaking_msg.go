package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreateValidator{}
	_ sdk.Msg = &MsgCreateValidator{}
)

func NewMsgCreateValidator(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	valAddr sdk.ValAddress, pubKey string,
	selfDelegation sdk.Coin, description Description, commission CommissionRates, minSelfDelegation math.Int,
) (*MsgCreateValidator, error) {
	selfDelegationConvert := Coin{
		Denom:  selfDelegation.Denom,
		Amount: selfDelegation.Amount,
	}
	return &MsgCreateValidator{
		Description:       description,
		DelegatorAddress:  sdk.AccAddress(valAddr).String(),
		ValidatorAddress:  valAddr.String(),
		Pubkey:            pubKey,
		Value:             selfDelegationConvert,
		Commission:        commission,
		MinSelfDelegation: minSelfDelegation,
		Creator:           creator,
		Port:              port,
		ChannelID:         channelID,
		TimeoutTimestamp:  timeoutTimestamp,
	}, nil
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgCreateValidator) ValidateBasic() error {
	// note that unmarshaling from bech32 ensures both non-empty and valid
	delAddr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid delegator address: %s", err)
	}
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}
	if !sdk.AccAddress(valAddr).Equals(delAddr) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validator address is invalid")
	}

	if msg.Pubkey == "" {
		return ErrEmptyValidatorPubKey
	}

	if !msg.Value.IsValid() || !msg.Value.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid delegation amount")
	}

	if msg.Description == (Description{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty description")
	}

	if msg.Commission == (CommissionRates{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty commission")
	}

	if err := msg.Commission.Validate(); err != nil {
		return err
	}

	if !msg.MinSelfDelegation.IsPositive() {
		return sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			"minimum self delegation must be a positive integer",
		)
	}

	if msg.Value.Amount.LT(msg.MinSelfDelegation) {
		return ErrSelfDelegationBelowMinimum
	}

	return nil
}

func (msg MsgCreateValidator) GetSigners() []sdk.AccAddress {
	// delegator is first signer so delegator pays fees
	delegator, _ := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	addrs := []sdk.AccAddress{delegator}
	valAddr, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)

	valAccAddr := sdk.AccAddress(valAddr)
	if !delegator.Equals(valAccAddr) {
		addrs = append(addrs, valAccAddr)
	}

	return addrs
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgCreateValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func NewDescription(moniker, identity, website, securityContact, details string) Description {
	return Description{
		Moniker:         moniker,
		Identity:        identity,
		Website:         website,
		SecurityContact: securityContact,
		Details:         details,
	}
}
