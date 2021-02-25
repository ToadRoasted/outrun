package responses

import (
	"github.com/ToadRoasted/outrun/consts"
	"github.com/ToadRoasted/outrun/meta"
	"github.com/ToadRoasted/outrun/responses/responseobjs"
)

type BaseResponse struct {
	responseobjs.BaseInfo
	AssetsVersion     string `json:"assets_version"` // doesn't necessarily have to be a number
	ClientDataVersion string `json:"client_data_version"`
	DataVersion       string `json:"data_version"`
	InfoVersion       string `json:"info_version"`
	Version           string `json:"version"`
	OutrunVersion     string `json:"ORN_version"`
}

func NewBaseResponse(base responseobjs.BaseInfo) BaseResponse {
	return BaseResponse{
		base,
		"051",
		"2.0.3",
		"15",
		"017",
		"2.0.3",
		meta.Version,
	}
}

func NewBaseResponseV(base responseobjs.BaseInfo, gameVersion string) BaseResponse {
	return BaseResponse{
		base,
		consts.DataVersionForGameVersion[gameVersion],
		gameVersion,
		"15",
		"017",
		gameVersion,
		meta.Version,
	}
}
