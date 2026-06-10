package user_domain

import (
	"context"

	"github.com/buqiuwenda/gin-template/internal/domain/user_domain/entity"
)

type UserRepository interface {
	CreateUser(user *entity.UserEntity) error
	GetUser(id uint64) (*entity.UserEntity, error)
}

type UserDomain struct {
	repo UserRepository
	ctx  context.Context
}

func NewUserDomain(repo UserRepository, ctx context.Context) *UserDomain {
	return &UserDomain{repo: repo, ctx: ctx}
}

func (u *UserDomain) CreateUser(user *entity.UserEntity) error {
	return u.repo.CreateUser(user)
}

func (u *UserDomain) GetUser(id uint64) (*entity.UserEntity, error) {
	return u.repo.GetUser(id)
}
