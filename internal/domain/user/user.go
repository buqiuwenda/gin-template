package user

import "time"

// User 领域实体（与 api proto DTO 解耦）
type User struct {
	ID        uint64
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
