package v1

import (
	"strconv"

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
	strID := ctx.Query("id")
	if strID == "" {
		reply.Error(ctx, ec.ErrInvalidParam)
		return
	}
	cusID, err := strconv.Atoi(strID)
	if err != nil {
		reply.Error(ctx, ec.ErrInvalidParam)
		log.Errorf("fail to parse customer id. %v", err)
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
