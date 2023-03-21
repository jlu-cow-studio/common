package http_struct

type ReqBase struct {
	LogId string `json:"log_id"`
	Token string `json:"token"`
}

type ResBase struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
