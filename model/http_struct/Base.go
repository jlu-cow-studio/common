package http_struct

type ReqBase struct {
	LogId string `json:"logid"`
	Token string `json:"token"`
}

type ResBase struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type OnlyBaseReq struct {
	Base ReqBase `json:"base"`
}

type OnlyBaseRes struct {
	Base ResBase `json:"base"`
}

var InvalidRequest OnlyBaseRes = OnlyBaseRes{ResBase{
	Message: "invalide request!",
	Code:    "499",
}}
