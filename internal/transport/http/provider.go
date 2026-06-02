package http

import (
	transportuser "github.com/buqiuwenda/gin-template/internal/transport/http/v1/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	transportuser.NewHandler,
)
