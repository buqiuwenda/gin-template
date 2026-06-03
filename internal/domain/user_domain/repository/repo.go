package repository

import (
	"context"
	"errors"

	"github.com/buqiuwenda/gin-template/internal/data/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	if u == nil {
		return nil, errors.New("user is nil")
	}
	m := &models.UserModel{Username: u.Username, Email: u.Email}
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}
	return toDomain(m), nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uint64) (*domain.User, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	var m models.UserModel
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return toDomain(&m), nil
}

func toDomain(m *models.UserModel) *domain.User {
	return &domain.User{
		ID:        m.ID,
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
