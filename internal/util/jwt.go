package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/AsdGroup8/ASD_QRCodeCheckIn/conf"
	"github.com/dgrijalva/jwt-go"
)

var claimPool = sync.Pool{
	New: func() interface{} {
		return new(UserClaim)
	},
}

// Jwt related Define
const (
	_tokenExpireIn = time.Hour * 8
	_claimIssuer   = "eequalsmc2"
)

// UserClaim all users shares the same jwt.Claim
type UserClaim struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// Free ...
func (uc *UserClaim) Free() {
	claimPool.Put(uc)
}

// GenToken generate signed token for user
func GenToken(userID uint) (string, error) {
	c := claimPool.Get().(*UserClaim)
	c.UserID = userID
	c.StandardClaims.ExpiresAt = time.Now().Add(_tokenExpireIn).Unix()
	c.StandardClaims.Issuer = _claimIssuer
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	c.Free()
	return token.SignedString(conf.Secret)
}

// ParseToken parse token to jwt.Claim
func ParseToken(tokenStr string) (*UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return conf.Secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail to parse token. %v", err)
	}
	if claims, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token. %v", err)
}
