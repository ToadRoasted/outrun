package muxhandlers

import (
	"encoding/json"

	"github.com/ToadRoasted/outrun/emess"
	"github.com/ToadRoasted/outrun/helper"
	"github.com/ToadRoasted/outrun/requests"
	"github.com/ToadRoasted/outrun/responses"
	"github.com/ToadRoasted/outrun/status"
)

func SendApollo(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.NewBaseResponseV(baseInfo, request.Version)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
