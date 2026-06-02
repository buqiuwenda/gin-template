package application

import (
	appuser "github.com/buqiuwenda/gin-template/internal/application/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	appuser.NewService,
)
