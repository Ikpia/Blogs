package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePost{}, "blogs/CreatePost", nil)
	cdc.RegisterConcrete(&MsgUpdatePost{}, "blogs/UpdatePost", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "blogs/DeletePost", nil)
	cdc.RegisterConcrete(&MsgWriteComment{}, "blogs/WriteComment", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeletePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWriteComment{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
