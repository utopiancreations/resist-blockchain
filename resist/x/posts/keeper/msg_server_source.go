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

func (k msgServer) CreateSource(ctx context.Context, msg *types.MsgCreateSource) (*types.MsgCreateSourceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Source.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var source = types.Source{
		Creator:          msg.Creator,
		Index:            msg.Index,
		Url:              msg.Url,
		Title:            msg.Title,
		Description:      msg.Description,
		CredibilityScore: msg.CredibilityScore,
		AnalysisSummary:  msg.AnalysisSummary,
		Verified:         msg.Verified,
	}

	if err := k.Source.Set(ctx, source.Index, source); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateSourceResponse{}, nil
}

func (k msgServer) UpdateSource(ctx context.Context, msg *types.MsgUpdateSource) (*types.MsgUpdateSourceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Source.Get(ctx, msg.Index)
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

	var source = types.Source{
		Creator:          msg.Creator,
		Index:            msg.Index,
		Url:              msg.Url,
		Title:            msg.Title,
		Description:      msg.Description,
		CredibilityScore: msg.CredibilityScore,
		AnalysisSummary:  msg.AnalysisSummary,
		Verified:         msg.Verified,
	}

	if err := k.Source.Set(ctx, source.Index, source); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update source")
	}

	return &types.MsgUpdateSourceResponse{}, nil
}

func (k msgServer) DeleteSource(ctx context.Context, msg *types.MsgDeleteSource) (*types.MsgDeleteSourceResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Source.Get(ctx, msg.Index)
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

	if err := k.Source.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove source")
	}

	return &types.MsgDeleteSourceResponse{}, nil
}
