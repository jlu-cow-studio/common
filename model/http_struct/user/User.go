package user

import "github.com/jlu-cow-studio/common/model/http_struct"

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
