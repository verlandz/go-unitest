package utils

import (
	"encoding/json"
	"net/http"
)

type HttpResp struct {
	Errors []string    `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// BuildHttpResp bulid response for http.
func BuildHttpResp(data interface{}, errs []error) HttpResp {
	_errs := []string{}
	for _, err := range errs {
		if err != nil {
			_errs = append(_errs, err.Error())
		}
	}

	return HttpResp{
		Errors: _errs,
		Data:   data,
	}
}

// JSON will build http reponse with
// 	"Header" Access-Control-Allow-Origin:*, Content-Type:application/json
func (resp HttpResp) JSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(b)
	return
}
