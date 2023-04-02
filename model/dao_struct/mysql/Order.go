package mysql

import "time"

type Order struct {
	OrderID          int       `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserID           int       `gorm:"column:user_id"`
	ItemID           int       `gorm:"column:item_id"`
	PurchaseDate     time.Time `gorm:"column:purchase_date"`
	PurchaseQuantity int       `gorm:"column:purchase_quantity"`
	UnitPrice        float64   `gorm:"column:unit_price"`
	TotalPrice       float64   `gorm:"column:total_price"`
}

type OrderForList struct {
	Order
	ItemInfo *Item `gorm:"foreignKey:item_id"`
}
