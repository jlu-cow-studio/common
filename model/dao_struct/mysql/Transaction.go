package mysql

import "time"

type Transaction struct {
	ID          int64     `gorm:"primary_key;column:id"`
	WalletID    int64     `gorm:"column:wallet_id"`
	Type        string    `gorm:"type:ENUM('deposit', 'purchase');column:type"`
	Amount      float64   `gorm:"type:decimal(10,2);column:amount"`
	OrderID     int64     `gorm:"column:order_id"`
	Description string    `gorm:"type:varchar(255);column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
