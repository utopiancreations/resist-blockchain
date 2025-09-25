package keeper

import (
	"context"
	"fmt"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateNodeResources(ctx context.Context, msg *types.MsgUpdateNodeResources) (*types.MsgUpdateNodeResourcesResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.NodeId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "node ID cannot be empty")
	}

	if err := ValidateResourceSpec(msg.AvailableResources); err != nil {
		return nil, errorsmod.Wrap(err, "invalid available resources")
	}

	// Get the node
	node, err := k.Nodes.Get(ctx, msg.NodeId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrNodeNotFound, "node not found")
	}

	// Check ownership
	if node.Owner != msg.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only node owner can update resources")
	}

	// Update node resources
	node.AvailableResources = msg.AvailableResources
	node.UptimePercentage = msg.UptimePercentage

	// Store updated node
	if err := k.Nodes.Set(ctx, msg.NodeId, node); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update node")
	}

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"node_resources_updated",
			sdk.NewAttribute("node_id", msg.NodeId),
			sdk.NewAttribute("owner", msg.Creator),
			sdk.NewAttribute("uptime_percentage", fmt.Sprintf("%d", msg.UptimePercentage)),
		),
	)

	return &types.MsgUpdateNodeResourcesResponse{}, nil
}