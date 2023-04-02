package trade

import (
	"github.com/jlu-cow-studio/common/model/http_struct"
	"github.com/jlu-cow-studio/common/model/http_struct/item"
)

type OrderInfo struct {
	ID        int64          `json:"id"`
	UserID    int64          `json:"user_id"`
	ItemID    int64          `json:"item_id"`
	Quantity  int32          `json:"quantity"`
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	ItemInfo  *item.ItemInfo `json:"item_info"`
}

type OrderListReq struct {
	Base    http_struct.ReqBase `json:"http_struct"`
	Page    int32               `json:"page"`
	PerPage int32               `json:"per_page"`
}

type OrderListRes struct {
	Base      http_struct.ResBase `json:"http_struct"`
	OrderList []*OrderInfo        `json:"order_list"`
}

type OrderReq struct {
	Base   http_struct.ReqBase `json:"http_struct"`
	ItemID int64               `json:"item_id"`
	Count  int32               `json:"count"`
}

type OrderRes struct {
	Base http_struct.ResBase `json:"http_struct"`
}

type RechargeReq struct {
	Base  http_struct.ReqBase `json:"http_struct"`
	Money float64             `json:"money"`
}

type RechargeRes struct {
	Base http_struct.ResBase `json:"http_struct"`
}
