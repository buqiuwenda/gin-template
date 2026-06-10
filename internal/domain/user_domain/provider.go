package user_domain

import (
	"github.com/buqiuwenda/gin-template/internal/domain/user_domain/repository"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserDomain,
	repository.NewUserRepository,
	wire.Bind(new(UserRepository), new(*repository.Repo)),
)
