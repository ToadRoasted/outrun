package responses

import (
    "github.com/ToadRoasted/outrun/netobj"
    "github.com/ToadRoasted/outrun/obj"
    "github.com/ToadRoasted/outrun/responses/responseobjs"
)

type RedStarExchangeListResponse struct {
    BaseResponse
    ItemList      []obj.RedStarItem `json:"itemList"`
    TotalItems    int64             `json:"totalItems"`
    MonthPurchase int64             `json:"monthPurchase"`
    Birthday      string            `json:"birthday"`
}

func RedStarExchangeList(base responseobjs.BaseInfo, itemList []obj.RedStarItem, monthPurchase int64, birthday string) RedStarExchangeListResponse {
    baseResponse := NewBaseResponse(base)
    totalItems := int64(len(itemList))
    return RedStarExchangeListResponse{
        baseResponse,
        itemList,
        totalItems,
        monthPurchase,
        birthday,
    }
}

func DefaultRedStarExchangeList(base responseobjs.BaseInfo) RedStarExchangeListResponse {
    itemList := []obj.RedStarItem{}
    monthPurchase := int64(0)
    birthday := "1900-1-1"
    return RedStarExchangeList(base, itemList, monthPurchase, birthday)
}

type RedStarExchangeResponse struct {
    BaseResponse
    PlayerState netobj.PlayerState `json:"playerState"`
}

func RedStarExchange(base responseobjs.BaseInfo, playerState netobj.PlayerState) RedStarExchangeResponse {
    baseResponse := NewBaseResponse(base)
    return RedStarExchangeResponse{
        baseResponse,
        playerState,
    }
}

func DefaultRedStarExchange(base responseobjs.BaseInfo, player netobj.Player) RedStarExchangeResponse {
    playerState := player.PlayerState
    return RedStarExchange(
        base,
        playerState,
    )
}
