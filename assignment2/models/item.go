package models

import "time"

type Item struct {
	ID          int       `gorm:"primaryKey;column:item_id;autoIncrement" json:"itemId,omitempty"`
	ItemCode    string    `gorm:"not null;column:item_code" json:"itemCode"`
	Description string    `gorm:"column:description" json:"description"`
	Quantity    int       `gorm:"column:quantity" json:"quantity"`
	OrderID     int       `gorm:"column:order_id" json:"orderId"`
	CreatedAt   time.Time `gorm:"type:TIMESTAMP(6);autoCreateTime"`
	UpdatedAt   time.Time `gorm:"type:TIMESTAMP(6);autoUpdateTime"`
}
