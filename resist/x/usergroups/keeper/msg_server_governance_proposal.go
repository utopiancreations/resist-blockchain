package keeper

import (
	"context"
	"errors"
	"fmt"

	"resist/x/usergroups/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateGovernanceProposal(ctx context.Context, msg *types.MsgCreateGovernanceProposal) (*types.MsgCreateGovernanceProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.GovernanceProposal.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var governanceProposal = types.GovernanceProposal{
		Creator:           msg.Creator,
		Index:             msg.Index,
		Title:             msg.Title,
		Description:       msg.Description,
		Proposer:          msg.Proposer,
		ProposalType:      msg.ProposalType,
		VotingPeriodStart: msg.VotingPeriodStart,
		VotingPeriodEnd:   msg.VotingPeriodEnd,
		YesVotes:          msg.YesVotes,
		NoVotes:           msg.NoVotes,
		AbstainVotes:      msg.AbstainVotes,
		Status:            msg.Status,
	}

	if err := k.GovernanceProposal.Set(ctx, governanceProposal.Index, governanceProposal); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateGovernanceProposalResponse{}, nil
}

func (k msgServer) UpdateGovernanceProposal(ctx context.Context, msg *types.MsgUpdateGovernanceProposal) (*types.MsgUpdateGovernanceProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.GovernanceProposal.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var governanceProposal = types.GovernanceProposal{
		Creator:           msg.Creator,
		Index:             msg.Index,
		Title:             msg.Title,
		Description:       msg.Description,
		Proposer:          msg.Proposer,
		ProposalType:      msg.ProposalType,
		VotingPeriodStart: msg.VotingPeriodStart,
		VotingPeriodEnd:   msg.VotingPeriodEnd,
		YesVotes:          msg.YesVotes,
		NoVotes:           msg.NoVotes,
		AbstainVotes:      msg.AbstainVotes,
		Status:            msg.Status,
	}

	if err := k.GovernanceProposal.Set(ctx, governanceProposal.Index, governanceProposal); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update governanceProposal")
	}

	return &types.MsgUpdateGovernanceProposalResponse{}, nil
}

func (k msgServer) DeleteGovernanceProposal(ctx context.Context, msg *types.MsgDeleteGovernanceProposal) (*types.MsgDeleteGovernanceProposalResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.GovernanceProposal.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.GovernanceProposal.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove governanceProposal")
	}

	return &types.MsgDeleteGovernanceProposalResponse{}, nil
}
