package domain

import (
	"github.com/buqiuwenda/gin-template/internal/domain/user_domain"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user_domain.ProviderSet,
)
