package mq_struct

import "github.com/jlu-cow-studio/common/model/dao_struct/redis"

type ItemChangeOp string

const (
	ItemOp_Create ItemChangeOp = "itemop_create"
	ItemOp_Update ItemChangeOp = "itemop_update"
	ItemOp_Delete ItemChangeOp = "itemop_delete"
)

type ItemChangeMsg struct {
	Op   ItemChangeOp `json:"op"`
	Info *redis.Item  `json:"item_info"`
}
