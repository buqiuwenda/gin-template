package recovery

import (
	"log"
	"net/http"

	"github.com/buqiuwenda/gin-template/internal/meta"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic recovered: %v", r)
				c.AbortWithStatusJSON(http.StatusInternalServerError, meta.Fail(500, "internal server error"))
			}
		}()
		c.Next()
	}
}
