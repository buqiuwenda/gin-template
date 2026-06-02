package jwt

import (
	"net/http"
	"strings"

	"github.com/buqiuwenda/gin-template/internal/config"
	"github.com/gin-gonic/gin"
)

// Auth JWT 鉴权中间件（模板占位，可按业务扩展 claims 校验）
func Auth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "unauthorized"})
			return
		}
		// TODO: 解析并校验 token
		_ = cfg.JWT.Secret
		c.Next()
	}
}
