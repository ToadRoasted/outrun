package muxhandlers

import (
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
)

func GetOptionUserResult(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.OptionUserResult(baseInfo, player.OptionUserResult)
	helper.SendResponse(response)
}
