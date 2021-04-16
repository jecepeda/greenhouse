package gtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/jecepeda/greenhouse/server/auth"
)

// TestRequest is a helper structure to simulate HTTP requests for tests
type TestRequest struct {
	Body         io.Reader
	DeviceID     uint64
	ExtraVars    map[string]string
	Method       string
	QueryVars    map[string]string
	ExtraHeaders map[string]string
	IsForm       bool
	IsRefresh    bool
}

func (tr TestRequest) Run(handler http.HandlerFunc) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(tr.Method, "", tr.Body)
	if tr.DeviceID != 0 {
		req = auth.WithAuth(req, tr.DeviceID, tr.IsRefresh)
	}
	if tr.IsForm {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	for k, v := range tr.ExtraHeaders {
		req.Header.Add(k, v)
	}
	vars := make(map[string]string)

	if tr.DeviceID != 0 {
		vars["deviceID"] = fmt.Sprint(tr.DeviceID)
	}

	for k, v := range tr.ExtraVars {
		vars[k] = v
	}

	q := req.URL.Query()
	for k, v := range tr.QueryVars {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	respRecord := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)

	handler.ServeHTTP(respRecord, req)
	return respRecord
}

func UnMarshalJSON(rr *httptest.ResponseRecorder, dest interface{}) error {
	data, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, dest)
	if err != nil {
		return err
	}
	return nil
}

func MarshalJSON(v interface{}) io.Reader {
	data, _ := json.Marshal(v)
	return bytes.NewReader(data)
}
