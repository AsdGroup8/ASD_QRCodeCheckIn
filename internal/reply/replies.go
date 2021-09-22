package reply

import (
	"net/http"
	"sync"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/gin-gonic/gin"
)

var resultPool = sync.Pool{
	New: func() interface{} { return new(Result) },
}

// OK ...
func OK(ctx *gin.Context, data interface{}) {
	result := resultPool.Get().(*Result)
	result.Code = int16(ec.Success)
	result.Msg = "Success"
	result.Data = data
	ctx.JSON(http.StatusOK, result)
	result.Data = nil
	resultPool.Put(result)
}

// Error ...
func Error(ctx *gin.Context, err error) {
	code := ec.Internal
	if svrErr, ok := err.(*ec.SvrErr); ok {
		code = svrErr.Code()
	}

	result := resultPool.Get().(*Result)
	result.Code = int16(code)
	result.Msg = err.Error()
	result.Data = nil
	ctx.JSON(http.StatusInternalServerError, result)
}
