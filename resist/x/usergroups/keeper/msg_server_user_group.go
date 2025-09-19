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

func (k msgServer) CreateUserGroup(ctx context.Context, msg *types.MsgCreateUserGroup) (*types.MsgCreateUserGroupResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.UserGroup.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var userGroup = types.UserGroup{
		Creator:       msg.Creator,
		Index:         msg.Index,
		Name:          msg.Name,
		Description:   msg.Description,
		Admin:         msg.Admin,
		Members:       msg.Members,
		VoteThreshold: msg.VoteThreshold,
		CreatedAt:     msg.CreatedAt,
	}

	if err := k.UserGroup.Set(ctx, userGroup.Index, userGroup); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateUserGroupResponse{}, nil
}

func (k msgServer) UpdateUserGroup(ctx context.Context, msg *types.MsgUpdateUserGroup) (*types.MsgUpdateUserGroupResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.UserGroup.Get(ctx, msg.Index)
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

	var userGroup = types.UserGroup{
		Creator:       msg.Creator,
		Index:         msg.Index,
		Name:          msg.Name,
		Description:   msg.Description,
		Admin:         msg.Admin,
		Members:       msg.Members,
		VoteThreshold: msg.VoteThreshold,
		CreatedAt:     msg.CreatedAt,
	}

	if err := k.UserGroup.Set(ctx, userGroup.Index, userGroup); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update userGroup")
	}

	return &types.MsgUpdateUserGroupResponse{}, nil
}

func (k msgServer) DeleteUserGroup(ctx context.Context, msg *types.MsgDeleteUserGroup) (*types.MsgDeleteUserGroupResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.UserGroup.Get(ctx, msg.Index)
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

	if err := k.UserGroup.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove userGroup")
	}

	return &types.MsgDeleteUserGroupResponse{}, nil
}
