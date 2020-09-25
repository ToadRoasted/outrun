package rpcobj

import (
	"github.com/Mtbcooler/outrun/logic"

	"github.com/Mtbcooler/outrun/db/dbaccess"
)

func (t *Toolbox) CalculateAndResetRankingData(nothing bool, reply *ToolboxReply) error {
	errcode, message := logic.CalculateAndResetRunnersLeague()
	if errcode == 1 {
		reply.Status = StatusOtherError
		reply.Info = message
		return nil
	}
	if errcode == 2 {
		reply.Status = StatusLeagueStillOngoing
		reply.Info = message
		return nil
	}
	if errcode == 3 {
		reply.Status = StatusLeagueNotStarted
		reply.Info = message
		return nil
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
