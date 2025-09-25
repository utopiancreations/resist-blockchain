package keeper

import (
	"context"
	"fmt"

	"resist/x/posts/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DistributeContent(ctx context.Context, msg *types.MsgDistributeContent) (*types.MsgDistributeContentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.ContentId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "content ID cannot be empty")
	}

	if len(msg.ContentData) == 0 {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "content data cannot be empty")
	}

	if msg.Metadata == nil {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "metadata cannot be nil")
	}

	if msg.TargetReplicas == 0 {
		msg.TargetReplicas = 3 // Default to 3 replicas
	}

	if msg.ReplicationStrategy == "" {
		msg.ReplicationStrategy = "geographic" // Default strategy
	}

	// Initialize content distribution service
	distributionService := NewContentDistributionService(&k.Keeper)

	// Distribute content to IPFS
	ipfsHash, err := distributionService.DistributeToIPFS(ctx, msg.ContentData, msg.Metadata)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to distribute content to IPFS")
	}

	// Select nodes for replication
	selectedNodes, err := distributionService.SelectReplicationNodes(
		ctx,
		msg.ReplicationStrategy,
		msg.TargetReplicas,
		msg.PreferredNodes,
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to select replication nodes")
	}

	// Generate distribution ID
	distributionId := fmt.Sprintf("dist_%s_%d", msg.ContentId, sdk.UnwrapSDKContext(ctx).BlockTime().Unix())

	// In a production environment, this would:
	// 1. Store content metadata in blockchain state
	// 2. Initiate replication to selected nodes
	// 3. Set up monitoring for replication status
	// 4. Handle payment for storage resources

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"content_distributed",
			sdk.NewAttribute("content_id", msg.ContentId),
			sdk.NewAttribute("ipfs_hash", ipfsHash),
			sdk.NewAttribute("distribution_id", distributionId),
			sdk.NewAttribute("target_replicas", fmt.Sprintf("%d", msg.TargetReplicas)),
			sdk.NewAttribute("replication_strategy", msg.ReplicationStrategy),
			sdk.NewAttribute("assigned_nodes", fmt.Sprintf("%v", selectedNodes)),
		),
	)

	return &types.MsgDistributeContentResponse{
		IpfsHash:       ipfsHash,
		AssignedNodes:  selectedNodes,
		DistributionId: distributionId,
	}, nil
}