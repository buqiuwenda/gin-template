//go:build wireinject
// +build wireinject

package web

import (
	"context"

	"github.com/buqiuwenda/gin-template/internal/application"
	"github.com/buqiuwenda/gin-template/internal/conf"
	"github.com/buqiuwenda/gin-template/internal/data"
	"github.com/buqiuwenda/gin-template/internal/domain"
	httpserver "github.com/buqiuwenda/gin-template/internal/server"
	transporthttp "github.com/buqiuwenda/gin-template/internal/transport/http"
	"github.com/google/wire"
)

func provideContext() context.Context {
	return context.Background()
}

func InitializeHTTPServer(bc *conf.Bootstrap) (*httpserver.HTTPServer, func(), error) {
	wire.Build(
		provideContext,
		wire.FieldsOf(new(*conf.Bootstrap), "Data"),
		data.ProviderSet,
		domain.ProviderSet,
		application.ProviderSet,
		transporthttp.ProviderSet,
		httpserver.ProviderSet,
	)
	return nil, nil, nil
}
