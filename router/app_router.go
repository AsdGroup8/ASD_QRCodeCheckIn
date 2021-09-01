package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitAppRouter ...
func InitAppRouter(r gin.IRouter) {
	apiApp := r.Group("/app")
	// TODO: JWT
	{
		apiApp.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})
	}
}
