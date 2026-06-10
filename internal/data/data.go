package data

import (
	"github.com/buqiuwenda/gin-template/internal/data/myclient"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	myclient.NewAigcMysqlClient,
)
