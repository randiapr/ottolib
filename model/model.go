package model

import "github.com/randiapr/ottolib/constant"

type PaginationMeta struct {
	TotalRecords int64 `json:"totalRecords"`
	TotalPages   int64 `json:"totalPages"`
	Offset       int64 `json:"offset"`
	Limit        int64 `json:"limit"`
	Page         int64 `json:"page"`
	PrevPage     int64 `json:"prevPage"`
	NextPage     int64 `json:"nextPage"`
}

// BaseResp: base response model
type BaseResp struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data"`
}

// BaseRespPage base response for pagination
type BaseRespPage struct {
	ResponseCode    string         `json:"responseCode"`
	ResponseMessage string         `json:"responseMessage"`
	Data            interface{}    `json:"data"`
	Pagination      PaginationMeta `json:"meta"`
}

func (br *BaseResp) OK(data interface{}) *BaseResp {
	return &BaseResp{ResponseCode: constant.RC_SUCCESS,
		ResponseMessage: "Success", Data: data}
}

func (brp *BaseRespPage) OK(data interface{}, meta PaginationMeta) *BaseRespPage {
	return &BaseRespPage{ResponseCode: constant.RC_SUCCESS,
		ResponseMessage: "Success", Data: data, Pagination: meta}
}
