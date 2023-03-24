package item

import "github.com/jlu-cow-studio/common/model/http_struct"

type AddItemReq struct {
	Base         http_struct.ReqBase `json:"base"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Category     string              `json:"category"`
	Price        float64             `json:"price"`
	Stock        int32               `json:"stock"`
	Province     string              `json:"province"`
	City         string              `json:"city"`
	District     string              `json:"district"`
	UserID       int32               `json:"user_id"`
	UserType     string              `json:"user_type"`
	SpecificAttr string              `json:"specific_attributes"`
}

type AddItemRes struct {
	Base   http_struct.ResBase `json:"base"`
	ItemID int32               `json:"item_id"`
}

type DeleteItemReq struct {
	Base   http_struct.ReqBase `json:"base"`
	ItemID int32               `json:"item_id"`
}

type DeleteItemRes struct {
	Base    http_struct.ResBase `json:"base"`
	Success bool                `json:"success"`
}
