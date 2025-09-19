package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGovernanceProposal{},
		&MsgUpdateGovernanceProposal{},
		&MsgDeleteGovernanceProposal{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateContentReport{},
		&MsgUpdateContentReport{},
		&MsgDeleteContentReport{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserGroup{},
		&MsgUpdateUserGroup{},
		&MsgDeleteUserGroup{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
