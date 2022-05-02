package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "market/CreatePool", nil)
	cdc.RegisterConcrete(&MsgJoinPool{}, "market/JoinPool", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "market/AddLiquidity", nil)
	cdc.RegisterConcrete(&MsgExitPool{}, "market/ExitPool", nil)
	cdc.RegisterConcrete(&MsgRemoveLiquidity{}, "market/RemoveLiquidity", nil)
	cdc.RegisterConcrete(&MsgSwap{}, "market/Swap", nil)
	cdc.RegisterConcrete(&MsgSendShares{}, "market/SendShares", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJoinPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExitPool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRemoveLiquidity{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSwap{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendShares{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
