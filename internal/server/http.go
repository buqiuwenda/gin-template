package server

import (
	"context"
	"net/http"
	"time"

	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/gin-gonic/gin"
)

// HTTPServer 封装 Gin HTTP 服务生命周期
type HTTPServer struct {
	cfg *config.Config
	srv *http.Server
}

func NewHTTPServer(cfg *config.Config, router *gin.Engine) *HTTPServer {
	return &HTTPServer{
		cfg: cfg,
		srv: &http.Server{
			Addr:    cfg.Server.Addr,
			Handler: router,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.srv.ListenAndServe()
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(shutdownCtx)
}
