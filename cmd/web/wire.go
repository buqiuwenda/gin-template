//go:build wireinject
// +build wireinject

package web

import (
	"github.com/buqiuwenda/gin-template/internal/app"
	"github.com/buqiuwenda/gin-template/internal/application"
	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/buqiuwenda/gin-template/internal/data"
	"github.com/buqiuwenda/gin-template/internal/server"
	transporthttp "github.com/buqiuwenda/gin-template/internal/transport/http"
	"github.com/google/wire"
)

func InitializeApp(configPath string) (*app.App, func(), error) {
	wire.Build(
		config.New,
		data.ProviderSet,
		application.ProviderSet,
		transporthttp.ProviderSet,
		server.ProviderSet,
		app.ProviderSet,
	)
	return nil, nil, nil
}
