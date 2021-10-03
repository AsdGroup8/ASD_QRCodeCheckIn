package v1

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/log"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	"github.com/gin-gonic/gin"
)

// OnGetAllMovie ...
func OnGetAllMovie(ctx *gin.Context) {
	movies, err := service.GetInTheatersMovies()
	if err != nil {
		log.Errorf("fail to get movies. %v", movies)
		reply.Error(ctx, ec.ErrInternal)
		return
	}
	reply.OK(ctx, movies)
}
