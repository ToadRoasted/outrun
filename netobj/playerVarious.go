package netobj

import (
	"github.com/Mtbcooler/outrun/config/gameconf"
)

// universal for all players

type PlayerVarious struct {
	CmSkipCount          int64 `json:"cmSkipCount"`
	EnergyRecoveryMax    int64 `json:"energyRecoveryMax"`
	EnergyRecoveryTime   int64 `json:"energyRecveryTime"`
	OnePlayCmCount       int64 `json:"onePlayCmCount"`
	OnePlayContinueCount int64 `json:"onePlayContinueCount"` // max. continues
	IsPurchased          int64 `json:"isPurchased"`
}

func DefaultPlayerVarious() PlayerVarious {
	cmSkipCount := int64(0)
	energyRecoveryMax := gameconf.CFile.EnergyRecoveryMax
	energyRecoveryTime := gameconf.CFile.EnergyRecoveryTime
	onePlayCmCount := int64(0)
	onePlayContinueCount := int64(2) // The cheat detection seems to trip if 3 or more continues are used.
	isPurchased := int64(0)
	return PlayerVarious{
		cmSkipCount,
		energyRecoveryMax,
		energyRecoveryTime,
		onePlayCmCount,
		onePlayContinueCount,
		isPurchased,
	}
}
