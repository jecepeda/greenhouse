package devicehandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jecepeda/greenhouse/server/auth"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/gerror"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/httputil"
)

const (
	tokenDuration   = 60 * time.Second
	refreshDuration = 24 * time.Hour
)

type loginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Login logins a device
func Login(dc handler.DependencyContainer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var (
			errMsg = "login"
		)
		ctx, cancel := context.WithTimeout(r.Context(), handler.DefaultDuration)
		defer cancel()
		r.ParseForm()

		strDevice := r.Form.Get("device")
		password := r.Form.Get("password")

		deviceID, err := strconv.ParseUint(strDevice, 10, 64)
		if err != nil {
			http.Error(rw, gerror.Wrap(err, errMsg).Error(), http.StatusBadRequest)
			return
		}

		d, err := dc.GetDeviceService().FindByID(ctx, deviceID)
		if err != nil {
			msg, code := device.ErrToHTTP(err)
			http.Error(rw, msg, code)
			return
		}

		err = dc.GetEncrypter().CheckPassword(d.Password, password)
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
		response := loginResponse{
			AccessToken:  authToken,
			RefreshToken: refreshToken,
		}
		httputil.WriteJSON(rw, response)
	}
}

type RefreshResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}

func Refresh(dc handler.DependencyContainer) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var errMsg = "refresh"

		claims, err := auth.GetJWTClaims(r)
		_, cancel := context.WithTimeout(r.Context(), handler.DefaultDuration)

		defer cancel()
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
		}

		httputil.WriteJSON(rw, RefreshResponse{AccessToken: accessToken})
	}
}
