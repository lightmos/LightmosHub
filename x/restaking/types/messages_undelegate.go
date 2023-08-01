package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"lightmos/types"
)

const TypeMsgSendUndelegate = "send_undelegate"

var _ sdk.Msg = &MsgSendUndelegate{}

func NewMsgSendUndelegate(
	validatorAddress string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	amount *types.Coin,
) *MsgSendUndelegate {
	return &MsgSendUndelegate{
		ValidatorAddress: validatorAddress,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Amount:           amount,
	}
}

func (msg *MsgSendUndelegate) Route() string {
	return RouterKey
}

func (msg *MsgSendUndelegate) Type() string {
	return TypeMsgSendUndelegate
}

func (msg *MsgSendUndelegate) GetSigners() []sdk.AccAddress {
	validatorAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{validatorAddress}
}

func (msg *MsgSendUndelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendUndelegate) ValidateBasic() error {
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
