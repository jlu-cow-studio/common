package tag

import "github.com/jlu-cow-studio/common/model/http_struct"

type Tag struct {
	ID         int64   `json:"tag_id"`
	Name       string  `json:"tag_name"`
	Weight     float64 `json:"weight"`
	MarkObject string  `json:"mark_object"`
	CategoryID int64   `json:"category_id"`
}

type TagWithCategory struct {
	ID         int64        `json:"tag_id"`
	Name       string       `json:"tag_name"`
	Weight     float64      `json:"weight"`
	MarkObject string       `json:"mark_object"`
	CategoryID int64        `json:"category_id"`
	Category   *TagCategory `json:"category"`
}

type TagCategory struct {
	ID       int    `json:"tag_category_id"`
	Name     string `json:"tag_category_name"`
	ParentID int    `json:"parent_id"`
	Level    int32  `json:"level"`
}

type TagCategoryWithList struct {
	TagCategory
	TagList []*Tag `json:"tag_list"`
}

type GetTagListBySceneRequest struct {
	Base  http_struct.ReqBase `json:"base"`
	Scene string              `json:"scene"`
}

type GetTagListBySceneResponse struct {
	Base http_struct.ResBase    `json:"base"`
	List []*TagCategoryWithList `json:"list"`
}

type GetTagListByItemReq struct {
	Base   http_struct.ReqBase `json:"base"`
	ItemID string              `json:"item_id"`
}

type GetTagListByItemRes struct {
	Base    http_struct.ResBase `json:"base"`
	TagList []*Tag              `json:"tag_list"`
}

type GetTagListByUserReq struct {
	Base http_struct.ReqBase `json:"base"`
}

type GetTagListByUserRes struct {
	Base    http_struct.ResBase `json:"base"`
	TagList []*Tag              `json:"tag_list"`
}

type UpdateUserTagsReq struct {
	Base    http_struct.ReqBase `json:"base"`
	TagList []string            `json:"tag_list"`
}

type UpdateUserTagsRes struct {
	Base http_struct.ResBase `jons:"base"`
}
