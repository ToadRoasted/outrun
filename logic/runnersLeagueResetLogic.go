package logic

import (
	"time"

	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/enums"
)

func CalculateAndResetRunnersLeague() (int, string) {
	_, leagueendtime, err := dbaccess.GetStartAndEndTimesForLeague(enums.RankingLeagueF_M, 0)
	if err == nil {
		if time.Now().UTC().Unix() <= leagueendtime {
			return 2, "NG: this command cannot be used while league is still ongoing"
		}
	}

	// TODO: Move people up and down leagues accordingly!

	err = dbaccess.ResetAllRankingLeagueData()
	if err != nil {
		return 1, "NG: unable to reset ranking data: " + err.Error()
	}

	err = dbaccess.ClearLeagueHighScores()
	if err != nil {
		return 1, "NG: unable to clear league high scores: " + err.Error()
	}
	return 0, "OK"
}
