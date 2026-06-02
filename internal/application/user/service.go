package user

import (
	"context"

	domain "github.com/buqiuwenda/gin-template/internal/domain/user"
)

// Service 应用层用例，编排 domain 与仓储
type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, username, email string) (*domain.User, error) {
	u := &domain.User{Username: username, Email: email}
	return s.repo.Create(ctx, u)
}

func (s *Service) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}
