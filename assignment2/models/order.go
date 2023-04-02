package models

import "time"

type Order struct {
	ID           int       `gorm:"primaryKey;column:order_id;autoIncrement" json:"orderId,omitempty"`
	CustomerName string    `gorm:"column:customer_name" json:"customerName"`
	OrderedAt    time.Time `gorm:"column:ordered_at" json:"orderedAt"`
	CreatedAt    time.Time `gorm:"type:TIMESTAMP(6);autoCreateTime"`
	UpdatedAt    time.Time `gorm:"type:TIMESTAMP(6);autoUpdateTime"`
	Items        []Item    `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
