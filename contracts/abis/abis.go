package abis

import (
	"github.com/umbracle/ethgo/abi"
)

var (
	// ABI for Staking Contract
	StakingABI = abi.MustNewABI(StakingJSONABI)

	// ABI for Contract used in e2e stress test
	StressTestABI = abi.MustNewABI(StressTestJSONABI)

	// ABI for Sample Contract used in e2e test
	SampleABI = abi.MustNewABI(SampleJSONABI)
)
