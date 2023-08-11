package types

import (
	"errors"
	"lightmos/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWithdrawToken = "withdraw_token"

var _ sdk.Msg = &MsgWithdrawToken{}

func NewMsgWithdrawToken(creator string, amount types.Coin) *MsgWithdrawToken {
	return &MsgWithdrawToken{
		Creator: creator,
		Value:   amount,
	}
}

func (msg *MsgWithdrawToken) Route() string {
	return RouterKey
}

func (msg *MsgWithdrawToken) Type() string {
	return TypeMsgWithdrawToken
}

func (msg *MsgWithdrawToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWithdrawToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWithdrawToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Value.Denom != "token" {
		return sdkerrors.Wrapf(errors.New("invalid voucher denom"), "invalid voucher denom (%s)", err)
	}
	return nil
}
