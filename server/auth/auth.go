package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jecepeda/greenhouse/server/runtime"
	"github.com/pkg/errors"
)

type JWTClaims struct {
	DeviceID uint64
	jwt.StandardClaims
}

func CreateJWT(deviceID uint64, refresh bool, duration time.Duration) (string, error) {
	var (
		tokenString string
		err         error
	)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS326"), &JWTClaims{
		deviceID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})

	if tokenString, err = token.SignedString(runtime.Vars.JWTSeedKey); err != nil {
		return "", errors.Wrap(err, "create jwt")
	}

	return tokenString, nil
}
