package v1

import (
	"strconv"

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
	r := &model.DBReservation{}
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
	packed := make([]*model.Reservation, 0, len(reservList))
	for i := range reservList {
		p := reservList[i].Packed()
		movie := service.GetMovieByID(reservList[i].MovieID)
		if movie != nil {
			p.MovieName = movie.Title
		} else {
			p.MovieName = "Unknown Movie"
		}
		packed = append(packed, p)
	}
	reply.OK(ctx, packed)
}

// OnCustomerDelReservHis ...
func OnCustomerDelReservHis(ctx *gin.Context) {
	strID := ctx.Query("id")
	if strID == "" {
		reply.Error(ctx, ec.ErrInvalidParam)
		return
	}
	reservID, err := strconv.Atoi(strID)
	if err != nil {
		reply.Error(ctx, ec.ErrInvalidParam)
		log.Errorf("fail to parse customer id. %v", err)
		return
	}
	if err := service.DeleteReservation(uint(reservID)); err != nil {
		reply.Error(ctx, ec.ErrInternal)
		return
	}
	reply.OK(ctx, nil)
}
