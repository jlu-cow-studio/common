package redis

type Item struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Category     string  `json:"category"`
	Price        float64 `json:"price"`
	Stock        int32   `json:"stock"`
	Province     string  `json:"province"`
	City         string  `json:"city"`
	District     string  `json:"district"`
	ImageURL     string  `json:"image_url"`
	UserID       int32   `json:"user_id"`
	UserType     string  `json:"user_type"`
	SpecificAttr string  `json:"specific_attributes"`
}

type ItemForFeed struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Category           string  `json:"category"`
	Price              float64 `json:"price"`
	Stock              int     `json:"stock"`
	Province           string  `json:"province"`
	City               string  `json:"city"`
	District           string  `json:"district"`
	ImageURL           string  `json:"image_url"`
	UserID             int     `json:"user_id"`
	UserType           string  `json:"user_type"`
	SpecificAttributes string  `json:"specific_attributes"`
	UID                int     `json:"uid"`
	Username           string  `json:"username"`
	UProvince          string  `json:"uprovince"`
	UCity              string  `json:"ucity"`
	UDistrict          string  `json:"udistrict"`
	URole              string  `json:"urole"`
}
