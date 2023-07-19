package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "lightmos/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendRetireShare = "send_retire_share"

var _ sdk.Msg = &MsgSendRetireShare{}

func NewMsgSendRetireShare(
	validatorAddress string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	amount *types.Coin,
) *MsgSendRetireShare {
	return &MsgSendRetireShare{
		ValidatorAddress: validatorAddress,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Amount:           amount,
	}
}

func (msg *MsgSendRetireShare) Route() string {
	return RouterKey
}

func (msg *MsgSendRetireShare) Type() string {
	return TypeMsgSendRetireShare
}

func (msg *MsgSendRetireShare) GetSigners() []sdk.AccAddress {
	validatorAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validatorAddress}
}

func (msg *MsgSendRetireShare) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendRetireShare) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validatorAddress address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
