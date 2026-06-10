package repository

import (
	"context"
	"errors"

	"github.com/buqiuwenda/gin-template/internal/data/models"
	"github.com/buqiuwenda/gin-template/internal/domain/user_domain/entity"
	"github.com/buqiuwenda/gin-template/internal/meta"
)

type Repo struct {
	ctx context.Context
	db  meta.MysqlWendaAigcClient
}

func NewUserRepository(ctx context.Context, db meta.MysqlWendaAigcClient) *Repo {
	return &Repo{
		ctx: ctx,
		db:  db,
	}
}

func (r *Repo) CreateUser(u *entity.UserEntity) error {
	if u == nil {
		return errors.New("user is nil")
	}
	m := &models.UserModel{Username: u.Username, Email: u.Email}
	_, err := r.db.Table(m.TableName()).Insert(m)
	if err != nil {
		return err
	}
	created := toDomain(m)
	u.ID = created.ID
	u.CreatedAt = created.CreatedAt
	u.UpdatedAt = created.UpdatedAt
	return nil
}

func (r *Repo) GetUser(id uint64) (*entity.UserEntity, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}
	var m models.UserModel
	_, err := r.db.Table(m.TableName()).Where("id = ?", id).Get(&m)
	if err != nil {
		return nil, err
	}
	return toDomain(&m), nil
}

func toDomain(m *models.UserModel) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        m.ID,
		Username:  m.Username,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
