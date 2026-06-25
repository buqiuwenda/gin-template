package models

import (
	"time"
)

type User struct {
	Id            uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	Nickname      string    `xorm:"not null comment('昵称') unique VARCHAR(100)"`
	Phone         string    `xorm:"not null comment('手机号') unique VARCHAR(20)"`
	Email         string    `xorm:"not null default '' comment('邮箱') VARCHAR(255)"`
	Password      string    `xorm:"not null default '' comment('密码') VARCHAR(255)"`
	Status        uint      `xorm:"not null default 2 comment('状态 1-启用 2-禁用') UNSIGNED TINYINT"`
	LastLoginTime time.Time `xorm:"comment('最后登录时间') DATETIME"`
	CreatedAt     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index DATETIME"`
	UpdatedAt     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
	DeletedAt     time.Time `xorm:"comment('删除时间') DATETIME"`
}

func (*User) TableName() string {
	return "user"
}
