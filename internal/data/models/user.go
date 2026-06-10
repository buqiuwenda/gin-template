package models

import "time"

// UserModel XORM 持久化模型（与 domain.User 解耦）
type UserModel struct {
	ID            uint64    `xorm:"pk autoincr" json:"id"`
	Username      string    `xorm:"size:64;uniqueIndex;not null" json:"username"`
	Phone         string    `xorm:"size:20;uniqueIndex;not null" json:"phone"`
	Email         string    `xorm:"size:128;uniqueIndex;not null" json:"email"`
	Password      string    `xorm:"size:128;not null" json:"password"`
	Status        uint8     `xorm:"not null default 2" json:"status"`
	LastLoginTime time.Time `xorm:"column:last_login_time" json:"last_login_time"`
	CreatedAt     time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `xorm:"updated_at" json:"updated_at"`
	DeletedAt     time.Time `xorm:"deleted_at" json:"deleted_at"`
}

func (*UserModel) TableName() string {
	return "user"
}
