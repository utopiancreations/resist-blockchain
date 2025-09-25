package keeper

import (
	"crypto/rand"
	"fmt"
	"resist/x/rewards/types"
)

// GenerateID generates a unique ID for offers and allocations
func GenerateID(prefix string) (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%x", prefix, bytes), nil
}

// ValidateResourceSpec validates that a resource specification is valid
func ValidateResourceSpec(spec *types.ResourceSpec) error {
	if spec == nil {
		return types.ErrInvalidResourceSpec
	}
	if spec.CpuCores == 0 && spec.MemoryGb == 0 && spec.StorageGb == 0 && spec.BandwidthMbps == 0 {
		return types.ErrInvalidResourceSpec
	}
	return nil
}

// HasSufficientResources checks if node has sufficient available resources
func HasSufficientResources(available, allocated, requested *types.ResourceSpec) bool {
	if available == nil || requested == nil {
		return false
	}

	// Calculate remaining available resources
	remainingCpu := available.CpuCores
	remainingMemory := available.MemoryGb
	remainingStorage := available.StorageGb
	remainingBandwidth := available.BandwidthMbps

	if allocated != nil {
		remainingCpu -= allocated.CpuCores
		remainingMemory -= allocated.MemoryGb
		remainingStorage -= allocated.StorageGb
		remainingBandwidth -= allocated.BandwidthMbps
	}

	// Check if remaining resources can satisfy the request
	return remainingCpu >= requested.CpuCores &&
		remainingMemory >= requested.MemoryGb &&
		remainingStorage >= requested.StorageGb &&
		remainingBandwidth >= requested.BandwidthMbps
}

// AddResourceSpec adds two resource specifications
func AddResourceSpec(a, b *types.ResourceSpec) *types.ResourceSpec {
	if a == nil {
		a = &types.ResourceSpec{}
	}
	if b == nil {
		return a
	}

	return &types.ResourceSpec{
		CpuCores:      a.CpuCores + b.CpuCores,
		MemoryGb:      a.MemoryGb + b.MemoryGb,
		StorageGb:     a.StorageGb + b.StorageGb,
		BandwidthMbps: a.BandwidthMbps + b.BandwidthMbps,
	}
}

// SubtractResourceSpec subtracts resource specification b from a
func SubtractResourceSpec(a, b *types.ResourceSpec) *types.ResourceSpec {
	if a == nil {
		return &types.ResourceSpec{}
	}
	if b == nil {
		return a
	}

	result := &types.ResourceSpec{
		CpuCores:      a.CpuCores,
		MemoryGb:      a.MemoryGb,
		StorageGb:     a.StorageGb,
		BandwidthMbps: a.BandwidthMbps,
	}

	if result.CpuCores >= b.CpuCores {
		result.CpuCores -= b.CpuCores
	} else {
		result.CpuCores = 0
	}

	if result.MemoryGb >= b.MemoryGb {
		result.MemoryGb -= b.MemoryGb
	} else {
		result.MemoryGb = 0
	}

	if result.StorageGb >= b.StorageGb {
		result.StorageGb -= b.StorageGb
	} else {
		result.StorageGb = 0
	}

	if result.BandwidthMbps >= b.BandwidthMbps {
		result.BandwidthMbps -= b.BandwidthMbps
	} else {
		result.BandwidthMbps = 0
	}

	return result
}