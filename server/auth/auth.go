package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/runtime"
	"github.com/pkg/errors"
)

var (
	validationKeyGetter = func(token *jwt.Token) (interface{}, error) {
		return []byte(runtime.Vars.JWTSeedKey), nil
	}
	ErrNoAuthFound = errors.New("no authorization header found")
)

type JWTClaims struct {
	DeviceID  uint64 `json:"device_id"`
	IsRefresh bool   `json:"is_refresh"`
	jwt.StandardClaims
}

func CreateJWT(deviceID uint64, refresh bool, duration time.Duration) (string, error) {
	var (
		tokenString string
		err         error
	)

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &JWTClaims{
		deviceID,
		refresh,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).UTC().Unix(),
		},
	})

	if tokenString, err = token.SignedString([]byte(runtime.Vars.JWTSeedKey)); err != nil {
		return "", gerror.Wrap(err, "create jwt")
	}

	return tokenString, nil
}

func GetJWTClaims(r *http.Request) (*JWTClaims, error) {
	rawAuth := r.Header.Get("Authorization")
	if rawAuth == "" {
		return nil, ErrNoAuthFound
	}
	rawAuth = strings.ReplaceAll(rawAuth, "Bearer ", "")

	p := &jwt.Parser{SkipClaimsValidation: true}
	tkn, err := p.Parse(rawAuth, validationKeyGetter)
	if err != nil {
		return nil, err
	}
	claims := tkn.Claims.(jwt.MapClaims)
	return &JWTClaims{
		DeviceID:  uint64(claims["device_id"].(float64)),
		IsRefresh: claims["is_refresh"].(bool),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(claims["exp"].(float64)),
		},
	}, nil
}

func GetJWTString(r *http.Request) (string, error) {
	rawAuth := r.Header.Get("Authorization")
	if rawAuth == "" {
		return "", ErrNoAuthFound
	}
	rawAuth = strings.ReplaceAll(rawAuth, "Bearer ", "")
	return rawAuth, nil
}

func MatchDeviceID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var deviceID uint64
		var err error
		vars := mux.Vars(r)

		if deviceID, err = strconv.ParseUint(vars["deviceID"], 10, 64); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		claims, err := GetJWTClaims(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		if claims.DeviceID != deviceID {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	claims, err := GetJWTClaims(r)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	if claims.Valid() != nil || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		fmt.Println("claims are not valid", claims.Valid(), claims.VerifyExpiresAt(time.Now().Unix(), true))
		http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	next(rw, r)
}

// WithAuth generates a new Request with JWT
func WithAuth(req *http.Request, deviceID uint64, isRefresh bool) *http.Request {
	token, _ := CreateJWT(deviceID, isRefresh, 60*time.Second)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}
