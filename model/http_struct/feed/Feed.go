package feed

import (
	"github.com/jlu-cow-studio/common/model/http_struct"
	"github.com/jlu-cow-studio/common/model/http_struct/item"
	"github.com/jlu-cow-studio/common/model/http_struct/user"
)

const (
	FeedScene_WholeCattle    = "whole_cattle"
	FeedScene_Cattle_product = "cattle_product"
	FeedScene_Breeding       = "breeding"
	FeedScene_Service        = "service"
	FeedScene_ServiceProduct = "service_product"
	FeedScene_VetPost        = "vet_post"
	FeedScene_LiveRoom       = "live_room"
	FeedScene_HomePageMix    = "homepage_mix"
)

var Role_FeedScene = map[string]map[string]interface{}{
	user.RoleBreeder: {
		FeedScene_WholeCattle:    nil,
		FeedScene_Cattle_product: nil,
		FeedScene_Breeding:       nil,
		FeedScene_Service:        nil,
		FeedScene_ServiceProduct: nil,
		FeedScene_LiveRoom:       nil,
		FeedScene_VetPost:        nil,
	},
	user.RoleNormal: {
		FeedScene_WholeCattle:    nil,
		FeedScene_Cattle_product: nil,
		FeedScene_Breeding:       nil,
	},
	user.RoleServiceProvider: {
		FeedScene_WholeCattle:    nil,
		FeedScene_Cattle_product: nil,
		FeedScene_Breeding:       nil,
		FeedScene_Service:        nil,
		FeedScene_ServiceProduct: nil,
	},
	user.RoleVeterinary: {
		FeedScene_WholeCattle:    nil,
		FeedScene_Cattle_product: nil,
		FeedScene_Breeding:       nil,
		FeedScene_LiveRoom:       nil,
		FeedScene_VetPost:        nil,
	},
}

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
