package server

import (
	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/buqiuwenda/gin-template/internal/middleware/jwt"
	"github.com/buqiuwenda/gin-template/internal/middleware/recovery"
	transportuser "github.com/buqiuwenda/gin-template/internal/transport/http/v1/user"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, userHandler *transportuser.Handler) *gin.Engine {
	gin.SetMode(cfg.Server.Mode)
	r := gin.New()
	r.Use(gin.Logger(), recovery.Recovery())

	api := r.Group("/api/v1")
	userHandler.Register(api)

	auth := api.Group("")
	auth.Use(jwt.Auth(cfg))
	_ = auth // 需鉴权路由挂到 auth 下

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	return r
}
