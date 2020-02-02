package muxhandlers

import (
	"github.com/Mtbcooler/outrun/consts"
	"github.com/Mtbcooler/outrun/db/dbaccess"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
)

func GetOptionUserResult(helper *helper.Helper) {
	if !helper.CheckSession(true) {
		return
	}
	pid, err := helper.GetCallingPlayerID()
	if err != nil {
		helper.InternalErr("Error getting calling player ID", err)
		return
	}
	optionUserResult, err := dbaccess.GetOptionUserResult(consts.DBMySQLTableOptionUserResults, pid)
	if err != nil {
		helper.InternalErr("Error getting OptionUserResult data", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.OptionUserResult(baseInfo, optionUserResult)
	helper.SendResponse(response)
}
