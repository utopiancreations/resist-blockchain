package keeper

import (
	"context"
	"fmt"
	"time"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AllocateResources(ctx context.Context, msg *types.MsgAllocateResources) (*types.MsgAllocateResourcesResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.OfferId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "offer ID cannot be empty")
	}

	if msg.ContentId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "content ID cannot be empty")
	}

	if err := ValidateResourceSpec(msg.RequestedResources); err != nil {
		return nil, errorsmod.Wrap(err, "invalid requested resources")
	}

	if msg.DurationHours <= 0 {
		return nil, errorsmod.Wrap(types.ErrInvalidDuration, "duration must be greater than 0")
	}

	// Get the resource offer
	offer, err := k.ResourceOffers.Get(ctx, msg.OfferId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrOfferNotFound, "resource offer not found")
	}

	if !offer.IsActive {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "resource offer is not active")
	}

	// Get the node
	node, err := k.Nodes.Get(ctx, offer.NodeId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrNodeNotFound, "node not found")
	}

	// Check if node has sufficient resources
	if !HasSufficientResources(node.AvailableResources, node.AllocatedResources, msg.RequestedResources) {
		return nil, errorsmod.Wrap(types.ErrInsufficientResources, "insufficient resources available")
	}

	// Generate allocation ID
	allocationId, err := GenerateID("alloc")
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to generate allocation ID")
	}

	// Calculate total cost
	totalCost := offer.PricePerHour * uint64(msg.DurationHours)

	// Create resource allocation
	allocation := types.ResourceAllocation{
		AllocationId:       allocationId,
		ContentId:          msg.ContentId,
		NodeId:             offer.NodeId,
		Requester:          msg.Creator,
		RequestedResources: msg.RequestedResources,
		AllocatedResources: msg.RequestedResources, // For now, allocate exactly what's requested
		StartTime:          time.Now().Unix(),
		EndTime:            time.Now().Unix() + (msg.DurationHours * 3600), // Convert hours to seconds
		Status:             "active",
		CostPerHour:        offer.PricePerHour,
		TotalCost:          totalCost,
	}

	// Store the allocation
	if err := k.ResourceAllocations.Set(ctx, allocationId, allocation); err != nil {
		return nil, errorsmod.Wrap(err, "failed to store resource allocation")
	}

	// Update node allocated resources
	node.AllocatedResources = AddResourceSpec(node.AllocatedResources, msg.RequestedResources)
	if err := k.Nodes.Set(ctx, offer.NodeId, node); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update node allocated resources")
	}

	// Update hub metrics
	metrics, err := k.HubMetrics.Get(ctx, offer.NodeId)
	if err != nil {
		// Create new metrics if doesn't exist
		metrics = types.HubMetrics{
			NodeId:            offer.NodeId,
			TotalAllocations:  0,
			ActiveAllocations: 0,
			TotalRevenue:      0,
			UptimeSeconds:     0,
			DataServedGb:      0,
			LastUpdated:       time.Now().Unix(),
		}
	}

	metrics.TotalAllocations++
	metrics.ActiveAllocations++
	metrics.TotalRevenue += totalCost
	metrics.LastUpdated = time.Now().Unix()

	if err := k.HubMetrics.Set(ctx, offer.NodeId, metrics); err != nil {
		return nil, errorsmod.Wrap(err, "failed to update hub metrics")
	}

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"resources_allocated",
			sdk.NewAttribute("allocation_id", allocationId),
			sdk.NewAttribute("node_id", offer.NodeId),
			sdk.NewAttribute("requester", msg.Creator),
			sdk.NewAttribute("content_id", msg.ContentId),
			sdk.NewAttribute("total_cost", fmt.Sprintf("%d", totalCost)),
			sdk.NewAttribute("duration_hours", fmt.Sprintf("%d", msg.DurationHours)),
		),
	)

	return &types.MsgAllocateResourcesResponse{
		AllocationId: allocationId,
		TotalCost:    totalCost,
	}, nil
}