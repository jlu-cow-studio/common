package feed

import (
	"github.com/jlu-cow-studio/common/model/http_struct"
	"github.com/jlu-cow-studio/common/model/http_struct/item"
)

type GetFeedReq struct {
	Base     http_struct.ReqBase `json:"base"`
	Scene    string              `json:"scene"`
	Page     int32               `json:"page"`
	PageSize int32               `json:"page_size"`
}

type GetFeedRes struct {
	Base     http_struct.ResBase `json:"base"`
	Items    []*item.ItemInfo    `json:"items"`
	Total    int32               `json:"total"`
	Page     int32               `json:"page"`
	PageSize int32               `json:"page_size"`
}

type ResetFeedReq struct {
	Base  http_struct.ReqBase `json:"base"`
	Scene string              `json:"scene"`
}

type ResetFeedRes struct {
	Base http_struct.ResBase `json:"base"`
}
