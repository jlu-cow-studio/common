package redis

type UserInfo struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Province string `json:"province"`
	ImageURL string `json:"image_url"`
	City     string `json:"city"`
	District string `json:"district"`
	Role     string `json:"role"`
}
