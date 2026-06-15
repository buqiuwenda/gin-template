package user

import (
	"context"
	"time"

	domain "github.com/buqiuwenda/gin-template/internal/domain/user_domain"
	"github.com/buqiuwenda/gin-template/internal/domain/user_domain/entity"
)

// Service 应用层用例，编排 domain 与仓储
type Service struct {
	userDomain *domain.UserDomain
	ctx        context.Context
}

func NewService(userDomain *domain.UserDomain, ctx context.Context) *Service {
	return &Service{userDomain: userDomain, ctx: ctx}
}

func (s *Service) CreateUser(ctx context.Context, username, email, password string) (*entity.UserEntity, error) {
	u := &entity.UserEntity{
		Username:      username,
		Email:         email,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Password:      password,
		LastLoginTime: time.Now(),
	}
	err := s.userDomain.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetUser(ctx context.Context, id uint64) (*entity.UserEntity, error) {
	return s.userDomain.GetUser(id)
}
