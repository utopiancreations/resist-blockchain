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

func (k msgServer) CreateSocialPost(ctx context.Context, msg *types.MsgCreateSocialPost) (*types.MsgCreateSocialPostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.SocialPost.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var socialPost = types.SocialPost{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Title:     msg.Title,
		Content:   msg.Content,
		MediaUrl:  msg.MediaUrl,
		MediaType: msg.MediaType,
		GroupId:   msg.GroupId,
		Author:    msg.Author,
		Upvotes:   msg.Upvotes,
		Downvotes: msg.Downvotes,
		CreatedAt: msg.CreatedAt,
	}

	if err := k.SocialPost.Set(ctx, socialPost.Index, socialPost); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateSocialPostResponse{}, nil
}

func (k msgServer) UpdateSocialPost(ctx context.Context, msg *types.MsgUpdateSocialPost) (*types.MsgUpdateSocialPostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.SocialPost.Get(ctx, msg.Index)
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

	var socialPost = types.SocialPost{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Title:     msg.Title,
		Content:   msg.Content,
		MediaUrl:  msg.MediaUrl,
		MediaType: msg.MediaType,
		GroupId:   msg.GroupId,
		Author:    msg.Author,
		Upvotes:   msg.Upvotes,
		Downvotes: msg.Downvotes,
		CreatedAt: msg.CreatedAt,
	}

	if err := k.SocialPost.Set(ctx, socialPost.Index, socialPost); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update socialPost")
	}

	return &types.MsgUpdateSocialPostResponse{}, nil
}

func (k msgServer) DeleteSocialPost(ctx context.Context, msg *types.MsgDeleteSocialPost) (*types.MsgDeleteSocialPostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.SocialPost.Get(ctx, msg.Index)
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

	if err := k.SocialPost.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove socialPost")
	}

	return &types.MsgDeleteSocialPostResponse{}, nil
}
