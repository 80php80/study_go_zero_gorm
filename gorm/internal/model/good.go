package model

import "time"

type Good struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // 主键
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
}

func (g Good) TableName() string {
	return "goods"
}
