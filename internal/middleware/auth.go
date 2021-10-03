package middleware

import (
	"strings"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/ec"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/reply"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/util"
	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware ...
func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			reply.Error(c, ec.ErrUnauthorized)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			reply.Error(c, ec.ErrInvalidToken)
			c.Abort()
			return
		}
		claim, err := util.ParseToken(parts[1])
		if err != nil {
			reply.Error(c, ec.ErrInvalidToken)
			c.Abort()
			return
		}
		// set user id
		c.Set(conf.StrCustomerID, claim.UserID)
		claim.Free()
		c.Next()
	}
}
