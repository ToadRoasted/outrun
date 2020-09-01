package responses

import "github.com/Mtbcooler/outrun/responses/responseobjs"

type BaseResponse struct {
	responseobjs.BaseInfo
	AssetsVersion     string `json:"assets_version"`
	ClientDataVersion string `json:"client_data_version"`
	DataVersion       string `json:"data_version"`
	InfoVersion       string `json:"info_version"`
	Version           string `json:"version"`
}

func NewBaseResponse(base responseobjs.BaseInfo) BaseResponse {
	return BaseResponse{
		base,
		"050", // TODO: Make this a const or something specifiable in the config file
		"2.0.3",
		"15",
		"017", // TODO: Make this a const or something specifiable in the config file
		"2.0.3",
	}

}

type NextVersionResponse struct {
	responseobjs.BaseInfo
	NumRedRingsIOS        int64  `json:"numRedRingsIOS,string"` // UNCONFIRMED!
	NumRedRingsANDROID    int64  `json:"numRedRingsANDROID,string"`
	NumBuyRedRingsIOS     int64  `json:"numBuyRedRingsIOS,string"` // UNCONFIRMED!
	NumBuyRedRingsANDROID int64  `json:"numBuyRedRingsANDROID,string"`
	Username              string `json:"userName"`
	CloseMessageJP        string `json:"closeMessageJP"`
	CloseMessageEN        string `json:"closeMessageEN"`
	CloseURL              string `json:"closeUrl"`
}

func NewNextVersionResponse(base responseobjs.BaseInfo, numRedRings, numBuyRedRings int64, username, japaneseMessage, englishMessage, url string) NextVersionResponse {
	return NextVersionResponse{
		base,
		numRedRings,
		numRedRings,
		numBuyRedRings,
		numBuyRedRings,
		username,
		japaneseMessage,
		englishMessage,
		url,
	}
}
