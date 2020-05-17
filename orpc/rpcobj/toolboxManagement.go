package rpcobj

import (
	"time"

	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/enums"
)

func (t *Toolbox) CalculateAndResetRankingData(nothing bool, reply *ToolboxReply) error {
	_, leagueendtime, err := dbaccess.GetStartAndEndTimesForLeague(enums.RankingLeagueF_M, 0)
	if err == nil {
		if time.Now().UTC().Unix() <= leagueendtime {
			reply.Status = StatusLeagueStillOngoing
			reply.Info = "this command cannot be used while league is still ongoing; use ForceResetRankingData instead"
			return nil
		}
	}

	err = dbaccess.ResetAllRankingLeagueData()
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to reset ranking data: " + err.Error()
		return err
	}

	err = dbaccess.ClearLeagueHighScores()
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to clear league high scores: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) ForceResetRankingData(nothing bool, reply *ToolboxReply) error {
	err := dbaccess.ResetAllRankingLeagueData()
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to reset ranking data: " + err.Error()
		return err
	}
	err = dbaccess.ClearLeagueHighScores()
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to clear league high scores: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}
