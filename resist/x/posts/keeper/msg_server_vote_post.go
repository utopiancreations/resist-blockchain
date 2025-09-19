package keeper

import (
	"context"
	"fmt"
	"strconv"

	"resist/x/posts/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) VotePost(ctx context.Context, msg *types.MsgVotePost) (*types.MsgVotePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// Validate vote type
	if msg.VoteType != "upvote" && msg.VoteType != "downvote" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "vote type must be 'upvote' or 'downvote'")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Find the social post
	postIndex := strconv.FormatUint(msg.PostId, 10)
	post, err := k.SocialPost.Get(ctx, postIndex)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "post not found")
	}

	// Create unique vote key (voter_address:post_id)
	voteKey := fmt.Sprintf("%s:%d", msg.Creator, msg.PostId)

	// Check if user already voted
	existingVote, err := k.Vote.Get(ctx, voteKey)
	if err == nil {
		// User has already voted, update the vote and adjust counts
		oldVoteType := existingVote.VoteType
		if oldVoteType != msg.VoteType {
			// User is changing their vote
			if oldVoteType == "upvote" {
				post.Upvotes--
			} else {
				post.Downvotes--
			}

			if msg.VoteType == "upvote" {
				post.Upvotes++
			} else {
				post.Downvotes++
			}

			// Update the vote
			existingVote.VoteType = msg.VoteType
			existingVote.Timestamp = sdkCtx.BlockTime().Unix()
			if err := k.Vote.Set(ctx, voteKey, existingVote); err != nil {
				return nil, errorsmod.Wrap(err, "failed to update vote")
			}
		} else {
			// Same vote type, no change needed
			return &types.MsgVotePostResponse{}, nil
		}
	} else {
		// New vote
		if msg.VoteType == "upvote" {
			post.Upvotes++
		} else {
			post.Downvotes++
		}

		// Create new vote record
		newVote := types.Vote{
			Creator:      msg.Creator,
			Index:        voteKey,
			VoterAddress: msg.Creator,
			PostId:       msg.PostId,
			VoteType:     msg.VoteType,
			Timestamp:    sdkCtx.BlockTime().Unix(),
		}

		if err := k.Vote.Set(ctx, voteKey, newVote); err != nil {
			return nil, errorsmod.Wrap(err, "failed to create vote")
		}
	}

	// Update the post with new vote counts
	if err := k.SocialPost.Set(ctx, postIndex, post); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update post vote counts")
	}

	// Emit vote event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"post_voted",
			sdk.NewAttribute("voter", msg.Creator),
			sdk.NewAttribute("post_id", strconv.FormatUint(msg.PostId, 10)),
			sdk.NewAttribute("vote_type", msg.VoteType),
			sdk.NewAttribute("upvotes", strconv.FormatUint(post.Upvotes, 10)),
			sdk.NewAttribute("downvotes", strconv.FormatUint(post.Downvotes, 10)),
		),
	)

	return &types.MsgVotePostResponse{}, nil
}
