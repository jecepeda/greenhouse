package devicehandler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jecepeda/greenhouse/server/auth"
	"github.com/jecepeda/greenhouse/server/domain/device"
	"github.com/jecepeda/greenhouse/server/handler"
	"github.com/jecepeda/greenhouse/server/handler/httputil"
	"github.com/pkg/errors"
)

type loginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func login(dc handler.DependencyContainer) http.HandlerFunc {
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
			http.Error(rw, errors.Wrap(err, errMsg).Error(), http.StatusBadRequest)
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

		authToken, err := auth.CreateJWT(d.ID, false, 1*time.Minute)
		if err != nil {
			http.Error(rw, errors.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}
		refreshToken, err := auth.CreateJWT(d.ID, true, 24*time.Hour)
		if err != nil {
			http.Error(rw, errors.Wrap(err, errMsg).Error(), http.StatusInternalServerError)
			return
		}
		response := loginResponse{
			AccessToken:  authToken,
			RefreshToken: refreshToken,
		}
		httputil.WriteJSON(rw, response)
	}
}
