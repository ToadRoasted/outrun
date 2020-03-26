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
