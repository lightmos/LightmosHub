package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSendSellOrder{}, "restaking/SendSellOrder", nil)
	cdc.RegisterConcrete(&MsgSendBuyOrder{}, "restaking/SendBuyOrder", nil)
	cdc.RegisterConcrete(&MsgCancelSellOrder{}, "restaking/CancelSellOrder", nil)
	cdc.RegisterConcrete(&MsgCancelBuyOrder{}, "restaking/CancelBuyOrder", nil)
	cdc.RegisterConcrete(&MsgSendRetireShare{}, "restaking/SendRetireShare", nil)
	cdc.RegisterConcrete(&MsgSendUndelegate{}, "restaking/SendUndelegate", nil)
	cdc.RegisterConcrete(&MsgWithdrawToken{}, "restaking/WithdrawToken", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelSellOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelBuyOrder{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendRetireShare{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendUndelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWithdrawToken{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
