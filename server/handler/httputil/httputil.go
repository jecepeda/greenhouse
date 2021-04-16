package httputil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jecepeda/greenhouse/server/gerror"
)

type ErrParameterNotFound struct {
	ParameterName string
}

func (e *ErrParameterNotFound) Error() string {
	return fmt.Sprintf("parameter %q not found", e.ParameterName)
}

func GetIDFromURL(r *http.Request, name string, data *uint64) error {
	errMsg := "get id from url"
	vars := mux.Vars(r)

	value, ok := vars[name]
	if !ok {
		return gerror.Wrap(&ErrParameterNotFound{ParameterName: name}, errMsg)
	}

	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return gerror.Wrap(err, errMsg)
	}

	*data = uint64(parsed)
	return nil
}

func ReadJSON(r *http.Request, data interface{}) error {
	errMsg := "read json"
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return gerror.Wrap(err, errMsg)
	}
	err = json.Unmarshal(raw, data)
	if err != nil {
		return gerror.Wrap(err, errMsg)
	}
	return nil
}

func WriteJSON(rw http.ResponseWriter, data interface{}) {
	raw, err := json.Marshal(&data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(raw)
}
