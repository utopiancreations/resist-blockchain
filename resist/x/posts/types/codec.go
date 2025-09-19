package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePostTag{},
		&MsgUpdatePostTag{},
		&MsgDeletePostTag{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSource{},
		&MsgUpdateSource{},
		&MsgDeleteSource{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVote{},
		&MsgUpdateVote{},
		&MsgDeleteVote{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSocialPost{},
		&MsgUpdateSocialPost{},
		&MsgDeleteSocialPost{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVotePost{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
