package keeper

import (
	"context"
	"fmt"
	"time"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateResourceOffer(ctx context.Context, msg *types.MsgCreateResourceOffer) (*types.MsgCreateResourceOfferResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.NodeId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "node ID cannot be empty")
	}

	if err := ValidateResourceSpec(msg.OfferedResources); err != nil {
		return nil, errorsmod.Wrap(err, "invalid offered resources")
	}

	if msg.PricePerHour == 0 {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "price per hour must be greater than 0")
	}

	// Check if node exists and is owned by creator
	node, err := k.Nodes.Get(ctx, msg.NodeId)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrNodeNotFound, "node not found")
	}

	if node.Owner != msg.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "only node owner can create resource offers")
	}

	// Check if node has sufficient resources
	if !HasSufficientResources(node.AvailableResources, node.AllocatedResources, msg.OfferedResources) {
		return nil, errorsmod.Wrap(types.ErrInsufficientResources, "node does not have sufficient resources")
	}

	// Generate offer ID
	offerId, err := GenerateID("offer")
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to generate offer ID")
	}

	// Create resource offer
	offer := types.ResourceOffer{
		OfferId:            offerId,
		NodeId:             msg.NodeId,
		Owner:              msg.Creator,
		AvailableResources: msg.OfferedResources,
		PricePerHour:       msg.PricePerHour,
		ContentTypes:       msg.ContentTypes,
		Location:           node.Location,
		IsActive:           true,
		CreatedAt:          time.Now().Unix(),
	}

	// Store the offer
	if err := k.ResourceOffers.Set(ctx, offerId, offer); err != nil {
		return nil, errorsmod.Wrap(err, "failed to store resource offer")
	}

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"resource_offer_created",
			sdk.NewAttribute("offer_id", offerId),
			sdk.NewAttribute("node_id", msg.NodeId),
			sdk.NewAttribute("owner", msg.Creator),
			sdk.NewAttribute("price_per_hour", fmt.Sprintf("%d", msg.PricePerHour)),
		),
	)

	return &types.MsgCreateResourceOfferResponse{
		OfferId: offerId,
	}, nil
}