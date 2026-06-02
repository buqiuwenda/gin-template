package user

import "context"

// Repository 仓储接口，由 data 层实现
type Repository interface {
	Create(ctx context.Context, u *User) (*User, error)
	GetByID(ctx context.Context, id uint64) (*User, error)
}
