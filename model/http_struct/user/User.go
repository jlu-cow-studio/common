package user

import (
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
