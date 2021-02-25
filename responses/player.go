package responses

import (
    "github.com/ToadRoasted/outrun/netobj"
    "github.com/ToadRoasted/outrun/responses/responseobjs"
)

type PlayerStateResponse struct {
    BaseResponse
    PlayerState netobj.PlayerState `json:"playerState"`
}

func PlayerState(base responseobjs.BaseInfo, playerState netobj.PlayerState) PlayerStateResponse {
    baseResponse := NewBaseResponse(base)
    out := PlayerStateResponse{
        baseResponse,
        playerState,
    }
    return out
}

type CharacterStateResponse struct {
    BaseResponse
    CharacterState []netobj.Character `json:"characterState"`
}

func CharacterState(base responseobjs.BaseInfo, characterState []netobj.Character) CharacterStateResponse {
    baseResponse := NewBaseResponse(base)
    out := CharacterStateResponse{
        baseResponse,
        characterState,
    }
    return out
}

type ChaoStateResponse struct {
    BaseResponse
    ChaoState []netobj.Chao `json:"chaoState"`
}

func ChaoState(base responseobjs.BaseInfo, chaoState []netobj.Chao) ChaoStateResponse {
    baseResponse := NewBaseResponse(base)
    out := ChaoStateResponse{
        baseResponse,
        chaoState,
    }
    return out
}
