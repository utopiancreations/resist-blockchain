package keeper

import (
	"context"
	"errors"
	"fmt"

	"resist/x/posts/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateVote(ctx context.Context, msg *types.MsgCreateVote) (*types.MsgCreateVoteResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Vote.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var vote = types.Vote{
		Creator:      msg.Creator,
		Index:        msg.Index,
		VoterAddress: msg.VoterAddress,
		PostIndex:    msg.PostIndex,
		VoteType:     msg.VoteType,
		Timestamp:    msg.Timestamp,
	}

	if err := k.Vote.Set(ctx, vote.Index, vote); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateVoteResponse{}, nil
}

func (k msgServer) UpdateVote(ctx context.Context, msg *types.MsgUpdateVote) (*types.MsgUpdateVoteResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Vote.Get(ctx, msg.Index)
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

	var vote = types.Vote{
		Creator:      msg.Creator,
		Index:        msg.Index,
		VoterAddress: msg.VoterAddress,
		PostIndex:    msg.PostIndex,
		VoteType:     msg.VoteType,
		Timestamp:    msg.Timestamp,
	}

	if err := k.Vote.Set(ctx, vote.Index, vote); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update vote")
	}

	return &types.MsgUpdateVoteResponse{}, nil
}

func (k msgServer) DeleteVote(ctx context.Context, msg *types.MsgDeleteVote) (*types.MsgDeleteVoteResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Vote.Get(ctx, msg.Index)
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

	if err := k.Vote.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove vote")
	}

	return &types.MsgDeleteVoteResponse{}, nil
}
