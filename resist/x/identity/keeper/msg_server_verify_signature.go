package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"resist/x/identity/types"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func (k msgServer) VerifySignature(ctx context.Context, msg *types.MsgVerifySignature) (*types.MsgVerifySignatureResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// Validate the address
	if _, err := k.addressCodec.StringToBytes(msg.Address); err != nil {
		return nil, errorsmod.Wrap(err, "invalid address")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := k.storeService.OpenKVStore(ctx)

	// Retrieve and validate challenge
	challengeKey := fmt.Sprintf("challenge:%s", msg.Address)
	challengeData, err := store.Get([]byte(challengeKey))
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to retrieve challenge")
	}
	if challengeData == nil {
		return nil, errorsmod.Wrap(types.ErrChallengeNotFound, "challenge not found for address")
	}

	// Parse challenge data (format: "challenge:expiration")
	parts := strings.Split(string(challengeData), ":")
	if len(parts) != 2 {
		return nil, errorsmod.Wrap(types.ErrInvalidChallenge, "invalid challenge format")
	}

	storedChallenge := parts[0]
	expiration, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidChallenge, "invalid expiration format")
	}

	// Check if challenge has expired
	if sdkCtx.BlockTime().Unix() > expiration {
		return nil, errorsmod.Wrap(types.ErrChallengeExpired, "challenge has expired")
	}

	// Verify the challenge matches
	if storedChallenge != msg.Challenge {
		return nil, errorsmod.Wrap(types.ErrInvalidChallenge, "challenge mismatch")
	}

	// Decode signature
	sigBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid signature format")
	}

	// Create message hash to verify
	messageToSign := fmt.Sprintf("Authenticate with challenge: %s", msg.Challenge)
	hash := sha256.Sum256([]byte(messageToSign))

	// Extract address bytes for verification
	_, addressBytes, err := bech32.DecodeAndConvert(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to decode address")
	}

	// Verify signature
	pubKey := &secp256k1.PubKey{Key: sigBytes[:33]} // First 33 bytes are public key
	signature := sigBytes[33:]                      // Remaining bytes are signature

	// Verify the signature
	if !pubKey.VerifySignature(hash[:], signature) {
		return nil, errorsmod.Wrap(types.ErrInvalidSignature, "signature verification failed")
	}

	// Verify the public key corresponds to the claimed address
	deriviedAddr := sdk.AccAddress(pubKey.Address())
	claimedAddr := sdk.AccAddress(addressBytes)
	if !deriviedAddr.Equals(claimedAddr) {
		return nil, errorsmod.Wrap(types.ErrAddressMismatch, "address does not match public key")
	}

	// Clean up the challenge
	if err := store.Delete([]byte(challengeKey)); err != nil {
		return nil, errorsmod.Wrap(err, "failed to delete challenge")
	}

	// Check if user profile exists, create/verify it
	profileExists, err := k.UserProfile.Has(ctx, msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to check profile existence")
	}

	if profileExists {
		// Update existing profile to mark as verified
		profile, err := k.UserProfile.Get(ctx, msg.Address)
		if err != nil {
			return nil, errorsmod.Wrap(err, "failed to get existing profile")
		}
		profile.Verified = true
		if err := k.UserProfile.Set(ctx, msg.Address, profile); err != nil {
			return nil, errorsmod.Wrap(err, "failed to verify profile")
		}
	} else {
		// Create new verified profile
		newProfile := types.UserProfile{
			Creator:     msg.Address,
			Index:       msg.Address,
			DisplayName: "", // User can set this later
			Bio:         "",
			AvatarUrl:   "",
			Verified:    true,
			CreatedAt:   sdkCtx.BlockTime().Unix(),
		}
		if err := k.UserProfile.Set(ctx, msg.Address, newProfile); err != nil {
			return nil, errorsmod.Wrap(err, "failed to create verified profile")
		}
	}

	// Emit successful authentication event
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			"authentication_successful",
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("creator", msg.Creator),
			sdk.NewAttribute("verified", "true"),
		),
	)

	return &types.MsgVerifySignatureResponse{}, nil
}
