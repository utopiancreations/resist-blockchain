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

func (k msgServer) CreatePostTag(ctx context.Context, msg *types.MsgCreatePostTag) (*types.MsgCreatePostTagResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.PostTag.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var postTag = types.PostTag{
		Creator:         msg.Creator,
		Index:           msg.Index,
		PostIndex:       msg.PostIndex,
		Tag:             msg.Tag,
		Category:        msg.Category,
		SimilarityScore: msg.SimilarityScore,
		RelatedPosts:    msg.RelatedPosts,
	}

	if err := k.PostTag.Set(ctx, postTag.Index, postTag); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreatePostTagResponse{}, nil
}

func (k msgServer) UpdatePostTag(ctx context.Context, msg *types.MsgUpdatePostTag) (*types.MsgUpdatePostTagResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.PostTag.Get(ctx, msg.Index)
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

	var postTag = types.PostTag{
		Creator:         msg.Creator,
		Index:           msg.Index,
		PostIndex:       msg.PostIndex,
		Tag:             msg.Tag,
		Category:        msg.Category,
		SimilarityScore: msg.SimilarityScore,
		RelatedPosts:    msg.RelatedPosts,
	}

	if err := k.PostTag.Set(ctx, postTag.Index, postTag); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update postTag")
	}

	return &types.MsgUpdatePostTagResponse{}, nil
}

func (k msgServer) DeletePostTag(ctx context.Context, msg *types.MsgDeletePostTag) (*types.MsgDeletePostTagResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.PostTag.Get(ctx, msg.Index)
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

	if err := k.PostTag.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove postTag")
	}

	return &types.MsgDeletePostTagResponse{}, nil
}
