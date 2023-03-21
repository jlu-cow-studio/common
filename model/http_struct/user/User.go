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
