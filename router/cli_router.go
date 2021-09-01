package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitCliRouter ...
func InitCliRouter(r gin.IRouter) {
	apiCli := r.Group("/cli")
	// TODO: JWT
	{
		apiCli.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "test")
		})
	}
}
