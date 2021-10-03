package v1

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	"github.com/gin-gonic/gin"
)

// OnCustomerCreateReserv ...
func OnCustomerCreateReserv(ctx *gin.Context) {
	cusID := ctx.GetUint(conf.StrCustomerID)
	if cusID == 0 {
		reply.Error(ctx, ec.ErrUnauthorized)
		return
	}
	r := &model.Reservation{}
	if err := ctx.ShouldBind(r); err != nil {
		log.Errorf("fail to bind. %v", err)
		reply.Error(ctx, ec.ErrInvalidParam)
		return
	}
	r.CustomerID = cusID
	if err := service.CreateReservation(r); err != nil {
		reply.Error(ctx, err)
		log.Errorf("fail to create reservation. %v", err)
		return
	}
	reply.OK(ctx, r)
}

// OnCustomerGetReservHis ...
func OnCustomerGetReservHis(ctx *gin.Context) {
	cusID := ctx.GetUint(conf.StrCustomerID)
	if cusID == 0 {
		reply.Error(ctx, ec.ErrUnauthorized)
		return
	}
	reservList, err := service.GetAllReservation(cusID)
	if err != nil {
		reply.Error(ctx, ec.ErrInternal)
		log.Errorf("fail to get customer %d's reservation.", cusID)
		return
	}
	reply.OK(ctx, reservList)
}

// OnCustomerDelReservHis ...
func OnCustomerDelReservHis(ctx *gin.Context) {
}
