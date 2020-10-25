package muxhandlers

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
	"github.com/jinzhu/now"
)

func GetPlayerState(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	//update energy counter
	for time.Now().UTC().Unix() >= player.PlayerState.EnergyRenewsAt && player.PlayerState.Energy < player.PlayerVarious.EnergyRecoveryMax {
		player.PlayerState.Energy++
		player.PlayerState.EnergyRenewsAt += player.PlayerVarious.EnergyRecoveryTime
	}
	if player.PlayerState.NextNumDailyChallenge <= 0 || int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
		// Initialize daily challenge if it isn't initialized already
		player.PlayerState.NumDailyChallenge = int64(0)
		player.PlayerState.NextNumDailyChallenge = int64(1)
		player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
	}
	if time.Now().UTC().Unix() >= player.PlayerState.DailyMissionEndTime {
		if player.PlayerState.DailyChallengeComplete == 1 && player.PlayerState.DailyChalSetNum < 10 {
			helper.DebugOut("Advancing to next daily mission...")
			player.PlayerState.DailyChalSetNum++
		} else {
			player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
			player.PlayerState.DailyChalSetNum = int64(0)
		}
		if player.PlayerState.DailyChallengeComplete == 0 {
			player.PlayerState.NumDailyChallenge = int64(0)
			player.PlayerState.NextNumDailyChallenge = int64(1)
		} else {
			player.PlayerState.NextNumDailyChallenge++
			if int(player.PlayerState.NextNumDailyChallenge) > len(consts.DailyMissionRewards) {
				player.PlayerState.NumDailyChallenge = int64(0)
				player.PlayerState.NextNumDailyChallenge = int64(1) //restart from beginning
				player.PlayerState.DailyChalCatNum = int64(rand.Intn(5))
				player.PlayerState.DailyChalSetNum = int64(0)
			}
		}
		player.PlayerState.DailyChalPosNum = int64(1 + rand.Intn(2))
		player.PlayerState.DailyMissionID = int64((player.PlayerState.DailyChalCatNum * 33) + (player.PlayerState.DailyChalSetNum * 3) + player.PlayerState.DailyChalPosNum)
		player.PlayerState.DailyChallengeValue = int64(0)
		player.PlayerState.DailyChallengeComplete = int64(0)
		player.PlayerState.DailyMissionEndTime = now.EndOfDay().UTC().Unix() + 1
		helper.DebugOut("New daily mission ID: %v", player.PlayerState.DailyMissionID)
	}
	player.PlayerState.LeagueStartTime, player.PlayerState.LeagueResetTime, err = dbaccess.GetStartAndEndTimesForLeague(player.PlayerState.RankingLeague, player.PlayerState.RankingLeagueGroup)
	if err != nil {
		helper.InternalErr("Error getting league start and end times (try restarting Outrun or sending the ResetRankingData RPC command)", err)
		return
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.PlayerState(baseInfo, player.PlayerState)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	helper.SendResponse(response)
}

func GetCharacterState(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.CharacterState(baseInfo, player.CharacterState)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	helper.SendResponse(response)
}

func GetChaoState(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.ChaoState(baseInfo, player.ChaoState)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	helper.SendResponse(response)
}

func SetUsername(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	recv := helper.GetGameRequest()
	var request requests.SetUsernameRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	// TODO: check if username is already taken
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	if len(request.Username) > 12 { // length checking, just in case someone hacks the name limit out of the game
		baseInfo.StatusCode = status.ClientError
	} else {
		player.Username = request.Username
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	response := responses.NewBaseResponse(baseInfo)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
