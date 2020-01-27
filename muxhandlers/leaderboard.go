package muxhandlers

import (
	"encoding/json"

	"github.com/Mtbcooler/outrun/consts"

	"github.com/Mtbcooler/outrun/db/dbaccess"

	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/obj/constobjs"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
)

func GetWeeklyLeaderboardOptions(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LeaderboardRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	uid, err := helper.GetCallingPlayerID()
	if err != nil {
		helper.InternalErr("Error getting calling player ID", err)
		return
	}
	playerState, err := dbaccess.GetPlayerState(consts.DBMySQLTablePlayerStates, uid)
	if err != nil {
		helper.InternalErr("Error getting player state", err)
		return
	}
	mode := request.Mode
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultWeeklyLeaderboardOptions(baseInfo, mode)
	response.StartTime = playerState.LeagueStartTime
	response.ResetTime = playerState.LeagueResetTime
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetWeeklyLeaderboardEntries(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LeaderboardEntriesRequest
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
	helper.DebugOut("Mode %v, type %v", request.Mode, request.Type)
	mode := request.Mode
	scoretype := request.Type
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultWeeklyLeaderboardEntries(baseInfo, player, mode, scoretype)
	response.StartTime = player.PlayerState.LeagueStartTime
	response.ResetTime = player.PlayerState.LeagueResetTime
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetLeagueData(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LeaderboardRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	uid, err := helper.GetCallingPlayerID()
	if err != nil {
		helper.InternalErr("Error getting calling player ID", err)
		return
	}
	playerState, err := dbaccess.GetPlayerState(consts.DBMySQLTablePlayerStates, uid)
	if err != nil {
		helper.InternalErr("Error getting player state", err)
		return
	}
	mode := request.Mode
	var leagueData obj.LeagueData
	if mode == 0 {
		leagueData = constobjs.LeagueDataDefinitions[playerState.RankingLeague]
		leagueData.GroupID = playerState.RankingLeagueGroup
	} else {
		leagueData = constobjs.LeagueDataDefinitions[playerState.QuickRankingLeague]
		leagueData.GroupID = playerState.QuickRankingLeagueGroup
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.LeagueData(baseInfo, leagueData, mode)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetLeagueOperatorData(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LeaderboardRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultLeagueOperatorData(baseInfo, 0)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
