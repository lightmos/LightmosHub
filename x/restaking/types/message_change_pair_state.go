package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangePairState = "change_pair_state"

var _ sdk.Msg = &MsgChangePairState{}

func NewMsgChangePairState(creator string, port string, channelID string, sourceDenom string, targetDenom string, price int32, author string, timeoutTimestamp uint64) *MsgChangePairState {
	return &MsgChangePairState{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		SourceDenom:      sourceDenom,
		TargetDenom:      targetDenom,
		Price:            price,
		Author:           author,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

func (msg *MsgChangePairState) Route() string {
	return RouterKey
}

func (msg *MsgChangePairState) Type() string {
	return TypeMsgChangePairState
}

func (msg *MsgChangePairState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangePairState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangePairState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
