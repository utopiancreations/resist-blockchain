package keeper

import (
	"context"
	"fmt"
	"time"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeallocateResources(ctx context.Context, msg *types.MsgDeallocateResources) (*types.MsgDeallocateResourcesResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.AllocationId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "allocation ID cannot be empty")
	}

	// Get the resource allocation
	allocation, err := k.ResourceAllocations.Get(ctx, msg.AllocationId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrAllocationNotFound, "resource allocation not found")
	}

	// Check authorization (only requester or node owner can deallocate)
	node, err := k.Nodes.Get(ctx, allocation.NodeId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrNodeNotFound, "node not found")
	}

	if allocation.Requester != msg.Creator && node.Owner != msg.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only requester or node owner can deallocate resources")
	}

	// Check if allocation is active
	if allocation.Status != "active" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "allocation is not active")
	}

	// Calculate refund amount (if deallocating early)
	var refundAmount uint64 = 0
	currentTime := time.Now().Unix()
	if currentTime < allocation.EndTime {
		// Calculate unused time
		unusedSeconds := allocation.EndTime - currentTime
		unusedHours := unusedSeconds / 3600
		if unusedSeconds%3600 > 0 {
			unusedHours++ // Round up
		}
		refundAmount = allocation.CostPerHour * uint64(unusedHours)

		// Ensure refund doesn't exceed total cost
		if refundAmount > allocation.TotalCost {
			refundAmount = allocation.TotalCost
		}
	}

	// Update allocation status
	allocation.Status = "completed"
	allocation.EndTime = currentTime // Update actual end time

	if err := k.ResourceAllocations.Set(ctx, msg.AllocationId, allocation); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update resource allocation")
	}

	// Update node allocated resources (subtract the deallocated resources)
	node.AllocatedResources = SubtractResourceSpec(node.AllocatedResources, allocation.AllocatedResources)
	if err := k.Nodes.Set(ctx, allocation.NodeId, node); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update node allocated resources")
	}

	// Update hub metrics
	metrics, err := k.HubMetrics.Get(ctx, allocation.NodeId)
	if err == nil {
		metrics.ActiveAllocations--
		if metrics.ActiveAllocations < 0 {
			metrics.ActiveAllocations = 0
		}
		metrics.LastUpdated = time.Now().Unix()

		if err := k.HubMetrics.Set(ctx, allocation.NodeId, metrics); err != nil {
			return nil, errorsmod.Wrap(err, "failed to update hub metrics")
		}
	}

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"resources_deallocated",
			sdk.NewAttribute("allocation_id", msg.AllocationId),
			sdk.NewAttribute("node_id", allocation.NodeId),
			sdk.NewAttribute("deallocator", msg.Creator),
			sdk.NewAttribute("refund_amount", fmt.Sprintf("%d", refundAmount)),
		),
	)

	return &types.MsgDeallocateResourcesResponse{
		RefundAmount: refundAmount,
	}, nil
}