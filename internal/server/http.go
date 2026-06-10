package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/buqiuwenda/gin-template/internal/conf"
	"github.com/buqiuwenda/seal/log"
	"github.com/gin-gonic/gin"
)

// HTTPServer 封装 Gin HTTP 服务生命周期（启动与优雅退出）
type HTTPServer struct {
	cfg *conf.Bootstrap
	srv *http.Server
}

func NewHTTPServer(cfg *conf.Bootstrap, router *gin.Engine) *HTTPServer {
	return &HTTPServer{
		cfg: cfg,
		srv: &http.Server{
			Addr:        cfg.Server.Http.Addr,
			ReadTimeout: cfg.Server.Http.Timeout.AsDuration(),
			Handler:     router,
		},
	}
}

func (s *HTTPServer) Run() error {
	go func() {
		log.Infof("http server listening on %s", s.cfg.Server.Http.Addr)
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.Server.Http.Timeout.AsDuration())
	defer cancel()
	return s.srv.Shutdown(ctx)
}
