package user

import (
	"github.com/jlu-cow-studio/common/model/dao_struct/redis"
	"github.com/jlu-cow-studio/common/model/http_struct"
)

const (
	RoleNormal          = "normal"
	RoleProducer        = "producer"
	RoleMerchant        = "merchant"
	RoleVeterinary      = "veterinary"
	RoleBreeder         = "breeder"
	RoleServiceProvider = "service_provider"

	FollowAction_Follow   = "follow"
	FollowAction_UnFollow = "unfollow"
)

type UserLoginReq struct {
	Base     http_struct.ReqBase `json:"base"`
	Username string              `json:"username"`
	Password string              `json:"password"`
}

type UserLoginRes struct {
	Base  http_struct.ResBase `json:"base"`
	Token string              `json:"token"`
}

type UserRegisterReq struct {
	Base     http_struct.ReqBase `json:"base"`
	Username string              `json:"username"`
	Password string              `json:"password"`
	Province string              `json:"province"`
	City     string              `json:"city"`
	District string              `json:"district"`
}

type UserRegisterRes struct {
	Base http_struct.ResBase `json:"base"`
}

type UserInfoReq struct {
	Base http_struct.ReqBase `json:"base"`
}

type UserInfoRes struct {
	Base     http_struct.ResBase `json:"base"`
	Username string              `json:"username"`
	Province string              `json:"province"`
	City     string              `json:"city"`
	District string              `json:"district"`
	Role     string              `json:"role"`
}

type UserAuthReq struct {
	Base http_struct.ReqBase `json:"base"`
	Role string              `json:"role"`
}

type UserAuthRes struct {
	Base http_struct.ResBase `json:"base"`
}

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
