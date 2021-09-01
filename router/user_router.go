package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitUserRouter ...
func InitUserRouter(r gin.IRouter) {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/TEST", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})
	}
}
