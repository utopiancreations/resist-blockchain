package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserProfile{},
		&MsgUpdateUserProfile{},
		&MsgDeleteUserProfile{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVerifySignature{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestChallenge{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
