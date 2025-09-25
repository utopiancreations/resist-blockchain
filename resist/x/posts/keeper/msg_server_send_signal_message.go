package keeper

import (
	"context"

	"resist/x/posts/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SendSignalMessage(ctx context.Context, msg *types.MsgSendSignalMessage) (*types.MsgSendSignalMessageResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid creator address")
	}

	// Validate input
	if msg.RecipientNode == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "recipient node cannot be empty")
	}

	if msg.ChannelId == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "channel ID cannot be empty")
	}

	if len(msg.EncryptedPayload) == 0 {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "encrypted payload cannot be empty")
	}

	if msg.MessageType == "" {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "message type cannot be empty")
	}

	// Validate message type
	validMessageTypes := map[string]bool{
		"resource_request": true,
		"resource_offer":   true,
		"sync_request":     true,
		"sync_response":    true,
		"content_alert":    true,
		"heartbeat":        true,
		"key_rotation":     true,
	}
	if !validMessageTypes[msg.MessageType] {
		return nil, errorsmod.Wrap(types.ErrInvalidInput, "invalid message type")
	}

	// Initialize Signal protocol service
	signalService := NewSignalProtocolService(&k.Keeper)

	// In production, this would validate the channel exists and sender has permission
	// For now, we'll assume the channel is valid

	// Generate message ID and send the message
	messageId, err := signalService.SendSecureMessage(
		ctx,
		msg.Creator, // sender node (derived from creator address)
		msg.RecipientNode,
		msg.ChannelId,
		msg.MessageType,
		msg.EncryptedPayload,
	)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to send Signal message")
	}

	// In a production environment, this would:
	// 1. Validate that the recipient node exists and is reachable
	// 2. Verify the sender has permission to send to this channel
	// 3. Queue the message for delivery if recipient is offline
	// 4. Implement delivery confirmations and retry logic
	// 5. Handle forward secrecy key rotation
	// 6. Rate limit messages to prevent spam

	// For demo purposes, assume delivery is always confirmed
	deliveryConfirmed := true

	// Emit event
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"signal_message_sent",
			sdk.NewAttribute("message_id", messageId),
			sdk.NewAttribute("sender", msg.Creator),
			sdk.NewAttribute("recipient_node", msg.RecipientNode),
			sdk.NewAttribute("channel_id", msg.ChannelId),
			sdk.NewAttribute("message_type", msg.MessageType),
			sdk.NewAttribute("delivery_confirmed", "true"),
		),
	)

	return &types.MsgSendSignalMessageResponse{
		MessageId:         messageId,
		DeliveryConfirmed: deliveryConfirmed,
	}, nil
}