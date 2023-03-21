package http_struct

type UserTokenInfo struct {
	Token    string `json:"token"`
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Role     int    `json:"role:"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}
