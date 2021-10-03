package router

import (
	apiv1 "github.com/AsdGroup8/ASD_QRCodeCheckIn/api/v1"
	"github.com/gin-gonic/gin"
)

// InitAppRouter ...
func InitAppRouter(r gin.IRouter) {
	app := r.Group("/app")
	{
		app.GET("/movie", apiv1.OnGetAllMovie)
	}
}
