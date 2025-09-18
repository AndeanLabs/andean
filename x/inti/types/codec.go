package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLazyBridgeTransfer{},
		&MsgUpdateLazyBridgeTransfer{},
		&MsgDeleteLazyBridgeTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitiateBridgeTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfirmBridgeTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateLazyTransfer{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
