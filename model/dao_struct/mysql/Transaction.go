package mysql

import "time"

type Transaction struct {
	ID          uint64    `gorm:"primary_key;column:id"`
	WalletID    uint64    `gorm:"column:wallet_id"`
	Type        string    `gorm:"type:ENUM('deposit', 'purchase');column:type"`
	Amount      float64   `gorm:"type:decimal(10,2);column:amount"`
	OrderID     uint64    `gorm:"column:order_id"`
	Description string    `gorm:"type:varchar(255);column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
