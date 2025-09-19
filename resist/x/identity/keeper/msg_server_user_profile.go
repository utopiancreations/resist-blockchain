package keeper

import (
	"context"
	"errors"
	"fmt"

	"resist/x/identity/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUserProfile(ctx context.Context, msg *types.MsgCreateUserProfile) (*types.MsgCreateUserProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Use the creator's address as the profile index for authentication integration
	profileIndex := msg.Creator

	// Check if the profile already exists
	ok, err := k.UserProfile.Has(ctx, profileIndex)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "user profile already exists")
	}

	// Set creation timestamp
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	currentTime := sdkCtx.BlockTime().Unix()

	var userProfile = types.UserProfile{
		Creator:     msg.Creator,
		Index:       profileIndex,
		DisplayName: msg.DisplayName,
		Bio:         msg.Bio,
		AvatarUrl:   msg.AvatarUrl,
		Verified:    false, // Profiles start as unverified
		CreatedAt:   currentTime,
	}

	if err := k.UserProfile.Set(ctx, userProfile.Index, userProfile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"profile_created",
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("display_name", msg.DisplayName),
		),
	)

	return &types.MsgCreateUserProfileResponse{}, nil
}

func (k msgServer) UpdateUserProfile(ctx context.Context, msg *types.MsgUpdateUserProfile) (*types.MsgUpdateUserProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.UserProfile.Get(ctx, msg.Index)
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

	var userProfile = types.UserProfile{
		Creator:     msg.Creator,
		Index:       msg.Index,
		DisplayName: msg.DisplayName,
		Bio:         msg.Bio,
		AvatarUrl:   msg.AvatarUrl,
		Verified:    msg.Verified,
		CreatedAt:   msg.CreatedAt,
	}

	if err := k.UserProfile.Set(ctx, userProfile.Index, userProfile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update userProfile")
	}

	return &types.MsgUpdateUserProfileResponse{}, nil
}

func (k msgServer) DeleteUserProfile(ctx context.Context, msg *types.MsgDeleteUserProfile) (*types.MsgDeleteUserProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.UserProfile.Get(ctx, msg.Index)
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

	if err := k.UserProfile.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove userProfile")
	}

	return &types.MsgDeleteUserProfileResponse{}, nil
}
