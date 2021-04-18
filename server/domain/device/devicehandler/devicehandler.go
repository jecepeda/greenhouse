package devicehandler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jecepeda/greenhouse/server/auth"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/httputil"
)

const (
	tokenDuration   = 60 * time.Second    // 60s (1 min)
	refreshDuration = 30 * 24 * time.Hour // 30 days 24h 3600s (1 hour)
)

type loginRequest struct {
	Device   uint64 `json:"device"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login logins a device
func Login(dc handler.DependencyContainer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var (
			errMsg       = "login"
			loginRequest loginRequest
		)
		ctx, cancel := context.WithTimeout(r.Context(), handler.DefaultDuration)
		defer cancel()

		if err := httputil.ReadJSON(r, &loginRequest); err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusBadRequest)
			return
		}

		d, err := dc.GetDeviceService().FindByID(ctx, loginRequest.Device)
		if err != nil {
			msg, code := device.ErrToHTTP(err)
			http.Error(rw, msg, code)
			return
		}

		err = dc.GetEncrypter().CheckPassword(d.Password, loginRequest.Password)
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		authToken, err := auth.CreateJWT(d.ID, false, tokenDuration)
		if err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}
		refreshToken, err := auth.CreateJWT(d.ID, true, refreshDuration)
		if err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(authToken)
		fmt.Println(refreshToken)
		response := loginResponse{
			AccessToken:  authToken,
			RefreshToken: refreshToken,
		}
		httputil.WriteJSON(rw, response)
	}
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Refresh refresh tokens for the given device, including the access token
// if its expiry date is out of date
func Refresh(dc handler.DependencyContainer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var errMsg = "refresh"
		_, cancel := context.WithTimeout(r.Context(), handler.DefaultDuration)
		defer cancel()

		claims, err := auth.GetJWTClaims(r)
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		if !claims.IsRefresh {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		accessToken, err := auth.CreateJWT(claims.DeviceID, false, tokenDuration)
		if err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}

		refreshToken, err := auth.GetJWTString(r)
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			// generate other token
			refreshToken, err = auth.CreateJWT(claims.DeviceID, true, refreshDuration)
			if err != nil {
				http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
				return
			}
		}

		httputil.WriteJSON(rw, RefreshResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}
}
