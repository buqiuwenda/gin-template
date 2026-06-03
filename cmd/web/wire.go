//go:build wireinject
// +build wireinject

package web

import (
	"github.com/buqiuwenda/gin-template/internal/application"
	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/buqiuwenda/gin-template/internal/data"
	"github.com/buqiuwenda/gin-template/internal/server"
	transporthttp "github.com/buqiuwenda/gin-template/internal/transport/http"
	"github.com/google/wire"
)

func InitializeHTTPServer(configPath string) (*server.HTTPServer, func(), error) {
	wire.Build(
		data.ProviderSet,
		application.ProviderSet,
		domain.ProviderSet,
		transporthttp.ProviderSet,
		server.ProviderSet,
	)
	return nil, nil, nil
}
