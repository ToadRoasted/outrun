package muxhandlers

import (
    "github.com/Mtbcooler/outrun/emess"
    "github.com/Mtbcooler/outrun/helper"
    "github.com/Mtbcooler/outrun/responses"
    "github.com/Mtbcooler/outrun/status"
)

func SendApollo(helper *helper.Helper) {
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.NewBaseResponse(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
