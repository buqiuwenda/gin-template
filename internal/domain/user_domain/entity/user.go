package entity

import "time"

// User 领域实体（与 api proto DTO 解耦）
type UserEntity struct {
	ID            uint64    `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Status        uint8     `json:"status"`
	LastLoginTime time.Time `json:"last_login_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
