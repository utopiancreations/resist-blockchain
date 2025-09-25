package keeper

import (
	"context"
	"time"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RegisterNode(ctx context.Context, msg *types.MsgRegisterNode) (*types.MsgRegisterNodeResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate required fields
	if msg.NodeId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "node ID cannot be empty")
	}
	if msg.NodeType == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "node type cannot be empty")
	}

	// Validate node type (could be "validator", "full", "light", etc.)
	validNodeTypes := map[string]bool{
		"validator": true,
		"full":      true,
		"light":     true,
		"archive":   true,
	}
	if !validNodeTypes[msg.NodeType] {
		return nil, errorsmod.Wrap(types.ErrInvalidNodeType, "node type must be one of: validator, full, light, archive")
	}

	// Check if node already exists
	_, err := k.Nodes.Get(ctx, msg.NodeId)
	if err == nil {
		return nil, errorsmod.Wrap(types.ErrNodeExists, "node with this ID already exists")
	}

	// Create the node
	node := types.Node{
		NodeId:                msg.NodeId,
		Owner:                 msg.Creator,
		NodeType:              msg.NodeType,
		StakeAmount:           msg.StakeAmount,
		CreatedAt:             time.Now().Unix(),
		IsActive:              true,
		AvailableResources:    msg.AvailableResources,
		AllocatedResources:    &types.ResourceSpec{}, // Initialize empty
		Location:              msg.Location,
		SupportedContentTypes: msg.SupportedContentTypes,
		UptimePercentage:      100, // Start at 100%
		BandwidthMbps:         msg.BandwidthMbps,
	}

	// Store the node
	if err := k.Nodes.Set(ctx, msg.NodeId, node); err != nil {
		return nil, errorsmod.Wrap(err, "failed to store node")
	}

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"node_registered",
			sdk.NewAttribute("node_id", msg.NodeId),
			sdk.NewAttribute("owner", msg.Creator),
			sdk.NewAttribute("node_type", msg.NodeType),
			sdk.NewAttribute("stake_amount", math.NewIntFromUint64(msg.StakeAmount).String()),
		),
	)

	return &types.MsgRegisterNodeResponse{}, nil
}
