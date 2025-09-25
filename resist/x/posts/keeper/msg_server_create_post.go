package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"resist/x/posts/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(ctx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate required fields
	if msg.Title == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "title cannot be empty")
	}
	if msg.Content == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "content cannot be empty")
	}

	// Generate unique index for the post
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	blockTime := sdkCtx.BlockTime().Unix()
	postIndex := fmt.Sprintf("%d-%d-%s", blockHeight, blockTime, msg.Creator)

	// Create the social post
	socialPost := types.SocialPost{
		Index:              postIndex,
		Title:              msg.Title,
		Content:            msg.Content,
		MediaUrl:           msg.MediaUrl,
		MediaType:          msg.MediaType,
		GroupId:            msg.GroupId,
		Author:             msg.Creator,
		Upvotes:            0,
		Downvotes:          0,
		CreatedAt:          uint64(time.Now().Unix()),
		Creator:            msg.Creator,
		Sources:            "[]", // Empty JSON array for sources
		Intent:             "discuss", // Default intent
		ContextType:        "opinion", // Default context type
		RequiresModeration: false,     // Default to not requiring moderation
	}

	// Store the social post
	if err := k.SocialPost.Set(ctx, postIndex, socialPost); err != nil {
		return nil, errorsmod.Wrap(err, "failed to store social post")
	}

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"post_created",
			sdk.NewAttribute("post_index", postIndex),
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("title", msg.Title),
			sdk.NewAttribute("group_id", strconv.FormatUint(msg.GroupId, 10)),
		),
	)

	return &types.MsgCreatePostResponse{}, nil
}
