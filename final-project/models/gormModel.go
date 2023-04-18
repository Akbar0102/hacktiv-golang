package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey;column:id;autoIncrement" json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
