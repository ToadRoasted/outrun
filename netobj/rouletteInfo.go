package netobj

import "github.com/jinzhu/now"

type RouletteInfo struct {
	LoginRouletteID       int64 `json:"ORN_loginRouletteId" db:"login_roulette_id"`
	RoulettePeriodEnd     int64 `json:"ORN_roulettePeriodEnd" db:"roulette_period_end"`
	RouletteCountInPeriod int64 `json:"ORN_rouletteCountInPeriod" db:"roulette_count_in_period"`
	GotJackpotThisPeriod  bool  `json:"ORN_gotJackpotThisPeriod" db:"got_jackpot_this_period"`
}

func DefaultRouletteInfo() RouletteInfo {
	loginRouletteID := int64(0)
	roulettePeriodEnd := now.EndOfDay().Unix()
	rouletteCountInPeriod := int64(0)
	gotJackpotThisPeriod := false
	return RouletteInfo{
		loginRouletteID,
		roulettePeriodEnd,
		rouletteCountInPeriod,
		gotJackpotThisPeriod,
	}
}
