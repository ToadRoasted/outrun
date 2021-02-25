package muxhandlers

import (
	"github.com/ToadRoasted/outrun/emess"
	"github.com/ToadRoasted/outrun/helper"
	"github.com/ToadRoasted/outrun/responses"
	"github.com/ToadRoasted/outrun/status"
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
