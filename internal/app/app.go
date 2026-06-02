package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/buqiuwenda/gin-template/internal/server"
)

// App 应用根，聚合 HTTP 服务
type App struct {
	http *server.HTTPServer
}

func New(httpSrv *server.HTTPServer) *App {
	return &App{http: httpSrv}
}

func (a *App) Run() error {
	go func() {
		log.Println("http server starting...")
		if err := a.http.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down...")
	return a.http.Stop(context.Background())
}
