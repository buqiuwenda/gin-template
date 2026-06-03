package user_domain


import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserDomain,
	repository.NewUserRepository,
)