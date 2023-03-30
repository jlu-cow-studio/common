package item

import "github.com/jlu-cow-studio/common/model/http_struct"

const (
	Category_WholeCattle    = "whole_cattle"
	Category_Cattle_product = "cattle_product"
	Category_Breeding       = "breeding"
	Category_Service        = "service"
	Category_ServiceProduct = "service_product"
)

type ItemInfo struct {
	ID                 int32   `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Category           string  `json:"category"`
	Price              float64 `json:"price"`
	Stock              int32   `json:"stock"`
	ImageURL           string  `json:"image_url"`
	Province           string  `json:"province"`
	City               string  `json:"city"`
	District           string  `json:"district"`
	UserID             int32   `json:"user_id"`
	UserType           string  `json:"user_type"`
	SpecificAttributes string  `json:"specific_attributes"`
}

type ItemForFeed struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Category           string  `json:"category"`
	Price              float64 `json:"price"`
	Stock              int     `json:"stock"`
	Province           string  `json:"province"`
	City               string  `json:"city"`
	District           string  `json:"district"`
	ImageURL           string  `json:"image_url"`
	UserID             int     `json:"user_id"`
	UserType           string  `json:"user_type"`
	SpecificAttributes string  `json:"specific_attributes"`
	UID                int     `json:"uid"`
	Username           string  `json:"username"`
	URole              string  `json:"urole"`
	UProvince          string  `json:"uprovince"`
	UCity              string  `json:"ucity"`
	UDistrict          string  `json:"udistrict"`
}

type AddItemReq struct {
	Base         http_struct.ReqBase `json:"base"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Category     string              `json:"category"`
	Price        float64             `json:"price"`
	Stock        int32               `json:"stock"`
	ImageUrl     string              `json:"image_url"`
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
	Base http_struct.ResBase `json:"base"`
}

type UpdateItemReq struct {
	Base         http_struct.ReqBase `json:"base"`
	ItemID       int32               `json:"item_id"`
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Price        float64             `json:"price"`
	Stock        int32               `json:"stock"`
	Province     string              `json:"province"`
	City         string              `json:"city"`
	District     string              `json:"district"`
	SpecificAttr string              `json:"specific_attributes"`
}

type UpdateItemRes struct {
	Base http_struct.ResBase `json:"base"`
}
