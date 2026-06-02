package user

import (
	"context"
	"errors"

	domain "github.com/buqiuwenda/gin-template/internal/domain/user"
)

type repo struct{}

func NewRepository() domain.Repository {
	return &repo{}
}

func (r *repo) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	_ = ctx
	if u == nil {
		return nil, errors.New("user is nil")
	}
	// TODO: 持久化到 MySQL
	u.ID = 1
	return u, nil
}

func (r *repo) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	_ = ctx
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	// TODO: 从 MySQL 查询
	return &domain.User{ID: id, Username: "demo", Email: "demo@example.com"}, nil
}
