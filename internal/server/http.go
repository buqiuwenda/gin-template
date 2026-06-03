package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/gin-gonic/gin"
)

// HTTPServer 封装 Gin HTTP 服务生命周期（启动与优雅退出）
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

func (s *HTTPServer) Run() error {
	go func() {
		log.Printf("http server listening on %s", s.cfg.Server.Addr)
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
