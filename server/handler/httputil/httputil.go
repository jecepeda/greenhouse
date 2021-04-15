package httputil

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(rw http.ResponseWriter, data interface{}) {
	raw, err := json.Marshal(&data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(raw)
}
