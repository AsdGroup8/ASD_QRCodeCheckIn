package v1

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/service"
	"github.com/gin-gonic/gin"
)

// OnGetAllMovie ...
func OnGetAllMovie(ctx *gin.Context) {
	reply.OK(ctx, service.GetInTheatersMovies())
}
