package user

import (
	"context"

	domain "github.com/buqiuwenda/gin-template/internal/domain/user_domain"
)

// Service 应用层用例，编排 domain 与仓储
type Service struct {
	userDomain user_domain.UserDomain
	ctx context.Context
}

func NewService(userDomain user_domain.UserDomain, ctx context.Context) *Service {
	return &Service{userDomain: userDomain, ctx: ctx}
}

func (s *Service) CreateUser(ctx context.Context, username, email string) (*domain.User, error) {
	u := &domain.User{Username: username, Email: email}
	return s.userDomain.CreateUser(ctx, u)
}

func (s *Service) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	return s.userDomain.GetUser(ctx, id)
}
