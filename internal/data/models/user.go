package models

import "time"

// UserModel XORM 持久化模型（与 domain.User 解耦）
type UserModel struct {
	ID            uint64    `xorm:"id" json:"id"`
	Nickname      string    `xorm:"nickname" json:"nickname"`
	Phone         string    `xorm:"phone" json:"phone"`
	Email         string    `xorm:"email" json:"email"`
	Password      string    `xorm:"password" json:"password"`
	Status        uint8     `xorm:"status" json:"status"`
	LastLoginTime time.Time `xorm:"last_login_time" json:"last_login_time"`
	CreatedAt     time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated_at" json:"updated_at"`
}

func (*UserModel) TableName() string {
	return "user"
}
