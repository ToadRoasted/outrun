package muxhandlers

import (
	"encoding/json"

	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/logic/conversion"
	"github.com/Mtbcooler/outrun/obj"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
	"github.com/jinzhu/now"
)

func GetDailyBattleData(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	response = responses.EmptyDailyBattleData(baseInfo)
	err := helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
}

func UpdateDailyBattleStatus(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	response = responses.UpdateDailyBattleStatus(baseInfo, now.EndOfDay().UTC().Unix(), obj.DefaultBattleStatus())

	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}

// Reroll daily battle rival
func ResetDailyBattleMatching(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.ResetDailyBattleMatchingRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	battleData := conversion.DebugPlayerToBattleData(player)
	startTime := now.BeginningOfDay().UTC().Unix()
	endTime := now.EndOfDay().UTC().Unix()
	var response interface{}
	response = responses.ResetDailyBattleMatchingNoOpponent(baseInfo, startTime, endTime, battleData, player)
	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
}

func GetDailyBattleHistory(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.GetDailyBattleHistoryRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	helper.DebugOut("Count: %v", request.Count)
	history := []obj.BattlePair{}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.GetDailyBattleHistory(baseInfo, history)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("error sending response", err)
	}
}

func GetDailyBattleStatus(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	battleStatus := obj.DefaultBattleStatus()
	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	response := responses.GetDailyBattleStatus(baseInfo, now.EndOfDay().UTC().Unix(), battleStatus)
	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}

func PostDailyBattleResult(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("error getting calling player", err)
		return
	}

	battleStatus := obj.DefaultBattleStatus()
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response interface{}
	response = responses.PostDailyBattleResultNoData(baseInfo,
		now.BeginningOfDay().UTC().Unix(),
		now.EndOfDay().UTC().Unix(),
		battleStatus,
	)

	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
}

func GetPrizeDailyBattle(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultGetPrizeDailyBattle(baseInfo)
	err := helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
