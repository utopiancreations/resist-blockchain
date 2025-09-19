package keeper

import (
	"context"

	"resist/x/rewards/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) RegisterNode(ctx context.Context, msg *types.MsgRegisterNode) (*types.MsgRegisterNodeResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgRegisterNodeResponse{}, nil
}
