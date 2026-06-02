package data

import (
	datauser "github.com/buqiuwenda/gin-template/internal/data/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	datauser.NewRepository,
)
