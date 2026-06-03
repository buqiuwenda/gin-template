package models

import "time"

// UserModel GORM 持久化模型（与 domain.User 解耦）
type UserModel struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"size:128;uniqueIndex;not null" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*UserModel) TableName() string {
	return "user"
}
