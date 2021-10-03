package router

import (
	apiv1 "github.com/AsdGroup8/ASD_QRCodeCheckIn/api/v1"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/middleware"
	"github.com/gin-gonic/gin"
)

// InitCustomerRouter ...
func InitCustomerRouter(r gin.IRouter) {
	api := r.Group("/customer")
	{
		api.POST("/reg", apiv1.OnCustomerRegister)
		api.POST("/auth", apiv1.OnCustomerAuth)
	}
	authReqapi := api.Use(middleware.UserAuthMiddleware())
	{
		authReqapi.GET("/profile", apiv1.OnCustomerGetProfile) // api => /customer/profile
		authReqapi.PUT("/profile", apiv1.OnCustomerEditProfile)

		authReqapi.GET("/reserv/history", apiv1.OnCustomerGetReservHis) // get all reservation history
		authReqapi.DELETE("/reserv/history", apiv1.OnCustomerDelReservHis)

		authReqapi.POST("/reserv", apiv1.OnCustomerCreateReserv) // api => /customer/reserv
	}
}
