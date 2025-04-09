package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"` // 主键
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (u User) TableName() string {
	return "users"
}
