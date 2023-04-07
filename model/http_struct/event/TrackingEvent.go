package event

import "github.com/jlu-cow-studio/common/model/http_struct"

const (
	Behavior_Browse         = "browse"
	Behavior_Click          = "click"
	Behavior_AddToFavorites = "add_to_favorites"
	Behavior_Purchase       = "purchase"
	Behavior_SearchBrowse   = "search_browse"
	Behavior_SearchClick    = "search_click"
	Behavior_LongView       = "long_view"
)

var BehaviorMap = map[string]interface{}{
	Behavior_Browse:         nil,
	Behavior_Click:          nil,
	Behavior_AddToFavorites: nil,
	Behavior_Purchase:       nil,
	Behavior_SearchBrowse:   nil,
	Behavior_SearchClick:    nil,
	Behavior_LongView:       nil,
}

type TrackingReportReq struct {
	Base     http_struct.ReqBase `json:"base"`
	ItemID   int32               `json:"item_id"`
	Behavior string              `json:"behavior"`
}

type TrackingReportRes struct {
	Base http_struct.ResBase `json:"base"`
}

type ViewingDurationReq struct {
	Base     http_struct.ReqBase `json:"base"`
	ItemID   int32               `json:"item_id"`
	Duration int64               `json:"duration"`
}

type ViewingDurationRes struct {
	Base http_struct.ResBase `json:"base"`
}
