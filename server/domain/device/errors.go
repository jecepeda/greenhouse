package device

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound = errors.New("device not found")
)

func ErrToHTTP(err error) (string, int) {
	if errors.Is(err, ErrNotFound) {
		return http.StatusText(http.StatusNotFound), http.StatusNotFound
	} else {
		return err.Error(), http.StatusInternalServerError
	}
}
