package router

import (
	apiv1 "github.com/AsdGroup8/ASD_QRCodeCheckIn/api/v1"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/middleware"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/gin-gonic/gin"
)

// InitCustomerRouter ...
func InitCustomerRouter(r gin.IRouter) {
	api := r.Group("/customer")
	{
		api.GET("/TEST", func(c *gin.Context) {
			reply.OK(c, "test")
		})
		api.POST("/reg", apiv1.OnCustomerRegister)
		api.POST("/auth", apiv1.OnCustomerAuth)
	}
	authReqapi := api.Use(middleware.UserAuthMiddleware())
	{
		authReqapi.GET("/profile", middleware.UserAuthMiddleware(), apiv1.OnCustomerGetProfile)
		authReqapi.PUT("/profile", middleware.UserAuthMiddleware(), apiv1.OnCustomerEditProfile)
	}
}
