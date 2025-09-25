package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/rewards module sentinel errors
var (
	ErrInvalidSigner        = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrInvalidInput         = errors.Register(ModuleName, 1101, "invalid input")
	ErrNodeExists           = errors.Register(ModuleName, 1102, "node already exists")
	ErrInvalidNodeType      = errors.Register(ModuleName, 1103, "invalid node type")
	ErrNodeNotFound         = errors.Register(ModuleName, 1104, "node not found")
	ErrOfferNotFound        = errors.Register(ModuleName, 1105, "resource offer not found")
	ErrAllocationNotFound   = errors.Register(ModuleName, 1106, "resource allocation not found")
	ErrInsufficientResources = errors.Register(ModuleName, 1107, "insufficient resources available")
	ErrInvalidDuration      = errors.Register(ModuleName, 1108, "invalid duration")
	ErrUnauthorized         = errors.Register(ModuleName, 1109, "unauthorized operation")
	ErrInvalidResourceSpec  = errors.Register(ModuleName, 1110, "invalid resource specification")
)
