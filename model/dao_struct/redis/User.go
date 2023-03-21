package redis

type UserInfo struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Role     int    `json:"role"`
}
