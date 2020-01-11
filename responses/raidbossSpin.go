package responses

import (
    "github.com/Mtbcooler/outrun/obj"
    "github.com/Mtbcooler/outrun/obj/constobjs"
    "github.com/Mtbcooler/outrun/responses/responseobjs"
)

type ItemStockNumResponse struct {
    BaseResponse
    ItemStockList []obj.Item `json:"itemStockList"`
}

func ItemStockNum(base responseobjs.BaseInfo, itemStockList []obj.Item) ItemStockNumResponse {
    baseResponse := NewBaseResponse(base)
    return ItemStockNumResponse{
        baseResponse,
        itemStockList,
    }
}

func DefaultItemStockNum(base responseobjs.BaseInfo) ItemStockNumResponse {
    return ItemStockNum(
        base,
        constobjs.DefaultSpinItems,
    )
}
