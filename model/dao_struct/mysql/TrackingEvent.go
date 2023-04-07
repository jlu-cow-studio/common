package mysql

import (
	"time"
)

type Event struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	ItemID    uint      `gorm:"column:item_id;not null" json:"item_id"`
	EventType string    `gorm:"column:event_type;not null" json:"event_type"`
	Timestamp time.Time `gorm:"column:timestamp;not null" json:"timestamp"`
	PageDepth *uint     `gorm:"column:page_depth" json:"page_depth,omitempty"`
	PageURL   *string   `gorm:"column:page_url" json:"page_url,omitempty"`
}

func (Event) TableName() string {
	return "events"
}
