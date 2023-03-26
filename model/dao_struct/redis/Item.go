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
	CreatedAt    int64   `json:"created_at"`
	SpecificAttr string  `json:"specific_attributes"`
}
