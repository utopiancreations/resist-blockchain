package keeper

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"resist/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestChallenge(ctx context.Context, msg *types.MsgRequestChallenge) (*types.MsgRequestChallengeResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// Validate the address to authenticate
	if _, err := k.addressCodec.StringToBytes(msg.Address); err != nil {
		return nil, errorsmod.Wrap(err, "invalid address to authenticate")
	}

	// Generate a random challenge
	challengeBytes := make([]byte, 32)
	if _, err := rand.Read(challengeBytes); err != nil {
		return nil, errorsmod.Wrap(err, "failed to generate challenge")
	}
	challengeString := hex.EncodeToString(challengeBytes)

	// Store challenge with expiration (5 minutes)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(ctx)
	challengeKey := fmt.Sprintf("challenge:%s", msg.Address)
	expiration := sdkCtx.BlockTime().Add(5 * time.Minute).Unix()
	challengeData := fmt.Sprintf("%s:%d", challengeString, expiration)

	if err := store.Set([]byte(challengeKey), []byte(challengeData)); err != nil {
		return nil, errorsmod.Wrap(err, "failed to store challenge")
	}

	// Emit event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"challenge_requested",
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("challenge", challengeString),
		),
	)

	return &types.MsgRequestChallengeResponse{}, nil
}
