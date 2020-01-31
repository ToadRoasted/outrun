package rpcobj

import (
	"github.com/Mtbcooler/outrun/db/dbaccess"
)

func (t *Toolbox) CalculateAndResetRankingData(nothing bool, reply *ToolboxReply) error {
	err := dbaccess.ResetRankingData()
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to reset ranking data: " + err.Error()
		return err
	}

	// TODO: Add league and reward logic!

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
