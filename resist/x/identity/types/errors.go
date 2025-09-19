package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/identity module sentinel errors
var (
	ErrInvalidSigner     = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrChallengeNotFound = errors.Register(ModuleName, 1101, "challenge not found")
	ErrInvalidChallenge  = errors.Register(ModuleName, 1102, "invalid challenge")
	ErrChallengeExpired  = errors.Register(ModuleName, 1103, "challenge has expired")
	ErrInvalidSignature  = errors.Register(ModuleName, 1104, "invalid signature")
	ErrAddressMismatch   = errors.Register(ModuleName, 1105, "address does not match public key")
)
