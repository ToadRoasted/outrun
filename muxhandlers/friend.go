package muxhandlers

import (
	"encoding/json"

	"github.com/Mtbcooler/outrun/db"
	"github.com/Mtbcooler/outrun/emess"
	"github.com/Mtbcooler/outrun/helper"
	"github.com/Mtbcooler/outrun/requests"
	"github.com/Mtbcooler/outrun/responses"
	"github.com/Mtbcooler/outrun/status"
)

func GetFacebookIncentive(helper *helper.Helper) {
	sid, _ := helper.GetSessionID()
	if !helper.CheckSession(true) {
		return
	}
	data := helper.GetGameRequest()
	var request requests.FacebookIncentiveRequest
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	switch request.Type {
	case 0:
		helper.DebugOut("Type 0 - LOGIN (facebook login)")
		break
	case 1:
		helper.DebugOut("Type 1 - REVIEW (review the game)")
		break
	case 2:
		helper.DebugOut("Type 2 - FEED (post to facebook feed)")
		break
	case 3:
		helper.DebugOut("Type 3 - ACHIEVEMENT (get achievement)")
		break
	case 4:
		helper.DebugOut("Type 4 - PUSH_NOLOGIN (respond to event push notif)")
		break
	default:
		helper.DebugOut("Unknown incentive type %v", request.Type)
		break
	}
	// We respond with no presents for now.
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultFacebookIncentive(baseInfo, player)
	response.Seq, _ = db.BoltGetSessionIDSeq(sid)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
