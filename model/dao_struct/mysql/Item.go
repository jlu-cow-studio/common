package mysql

import (
	"time"

	"github.com/jlu-cow-studio/common/model/dao_struct/redis"
)

// 定义商品分类的常量值
const (
	CategoryWholeCattle    = "whole_cattle"
	CategoryBreeding       = "breeding"
	CategoryCattleProduct  = "cattle_product"
	CategoryServiceProduct = "service_product"
	CategoryService        = "service"
)

type Item struct {
	ID           int32     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name         string    `gorm:"column:name;not null"`
	Description  string    `gorm:"column:description"`
	Category     string    `gorm:"column:category;type:enum('whole_cattle', 'breeding', 'cattle_product', 'service_product', 'service');not null"`
	Price        float64   `gorm:"column:price;type:decimal(10,2);not null"`
	Stock        int32     `gorm:"column:stock;default:0"`
	Province     string    `gorm:"column:province"`
	City         string    `gorm:"column:city"`
	District     string    `gorm:"column:district"`
	ImageURL     string    `gorm:"column:image_url"`
	UserID       int32     `gorm:"column:user_id;not null"`
	UserType     string    `gorm:"column:user_type;type:enum('breeder', 'service_provider');not null"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	SpecificAttr string    `gorm:"column:specific_attributes;type:json"`
}

// TableName sets the name of the table for the Item model.
func (Item) TableName() string {
	return "items"
}

func (i *Item) ToRedis() *redis.Item {
	return &redis.Item{
		ID:           i.ID,
		Name:         i.Name,
		Description:  i.Description,
		Category:     i.Category,
		Price:        i.Price,
		Stock:        i.Stock,
		Province:     i.Province,
		City:         i.City,
		District:     i.District,
		ImageURL:     i.ImageURL,
		UserID:       i.UserID,
		UserType:     i.UserType,
		CreatedAt:    i.CreatedAt.Unix(),
		SpecificAttr: i.SpecificAttr,
	}
}

func ItemFromRedis(redisItem *redis.Item) *Item {
	return &Item{
		ID:           redisItem.ID,
		Name:         redisItem.Name,
		Description:  redisItem.Description,
		Category:     redisItem.Category,
		Price:        redisItem.Price,
		Stock:        redisItem.Stock,
		Province:     redisItem.Province,
		City:         redisItem.City,
		District:     redisItem.District,
		ImageURL:     redisItem.ImageURL,
		UserID:       redisItem.UserID,
		UserType:     redisItem.UserType,
		CreatedAt:    time.Unix(redisItem.CreatedAt, 0),
		SpecificAttr: redisItem.SpecificAttr,
	}
}
