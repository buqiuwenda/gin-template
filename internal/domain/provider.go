package domain

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	user_domain.ProviderSet,
)