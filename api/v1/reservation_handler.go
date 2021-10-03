package v1

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	"github.com/gin-gonic/gin"
)

// OnCustomerCreateReserv ...
func OnCustomerCreateReserv(ctx *gin.Context) {
}

// OnCustomerGetReservHis ...
func OnCustomerGetReservHis(ctx *gin.Context) {
	cusID := ctx.GetUint(conf.StrUserID)
	if cusID == 0 {
		reply.Error(ctx, ec.ErrUnauthorized)
		return
	}
	reservList := service.GetAllReservation(cusID)
	if reservList == nil {
		reply.Error(ctx, ec.ErrInternal)
		log.Errorf("fail to get customer %d's reservation.", cusID)
		return
	}
	reply.OK(ctx, reservList)
}

// OnCustomerDelReservHis ...
func OnCustomerDelReservHis(ctx *gin.Context) {
}
