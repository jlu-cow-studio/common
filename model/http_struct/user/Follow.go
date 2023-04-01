package user

import (
	"github.com/jlu-cow-studio/common/model/dao_struct/redis"
	"github.com/jlu-cow-studio/common/model/http_struct"
)

type UserFollowReq struct {
	Base        http_struct.ReqBase `json:"base"`
	FollowingId string              `json:"following_id"`
	Action      string              `json:"action"`
}

type UserFollowRes struct {
	Base http_struct.ResBase `json:"base"`
}

type UserFollowingReq struct {
	Base     http_struct.ReqBase `json:"base"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

type UserFollowingRes struct {
	Base       http_struct.ResBase `json:"base"`
	TotalCount int                 `json:"total_count"`
	TotalPage  int                 `json:"total_page"`
	Users      []*redis.UserInfo   `json:"users"`
}

type UserFollowerReq struct {
	Base     http_struct.ReqBase `json:"base"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
}

type UserFollowerRes struct {
	Base       http_struct.ResBase `json:"base"`
	TotalCount int                 `json:"total_count"`
	TotalPage  int                 `json:"total_page"`
	Users      []*redis.UserInfo   `json:"users"`
}

type UserFollowCountReq struct {
	Base http_struct.ReqBase `json:"base"`
}

type UserFollowCountRes struct {
	Base           http_struct.ResBase `json:"base"`
	FollowingCount int                 `json:"following_count"`
	FollowerCount  int                 `json:"follower_count"`
}
