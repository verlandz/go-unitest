package http

import (
	_http "net/http"
	"net/http/httptest"
	"testing"
)

type TestGetRandNumber struct {
	N int
}

// FailClientNil fail because client is nil.
func (tc TestGetRandNumber) FailClientNil(t *testing.T) (actual int, err error) {
	r := New(nil)
	return r.GetRandNumber(tc.N)
}

// FailClientNewRequest fail to create NewRequest.
func (tc TestGetRandNumber) FailClientNewRequest(t *testing.T) (actual int, err error) {
	defer func(backupURL string) { url = backupURL }(url)
	url = "://"

	r := New(&_http.Client{})
	return r.GetRandNumber(tc.N)
}

// FailClientDo fail to send http request.
func (tc TestGetRandNumber) FailClientDo(t *testing.T) (actual int, err error) {
	defer func(backupURL string) { url = backupURL }(url)
	url = ""

	r := New(&_http.Client{})
	return r.GetRandNumber(tc.N)
}

// FailClientUnmarshall got "" from dependency.
func (tc TestGetRandNumber) FailClientUnmarshall(t *testing.T) (actual int, err error) {
	defer func(backupURL string) { url = backupURL }(url)

	mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
		func(writer _http.ResponseWriter, request *_http.Request) {
			writer.Write([]byte(""))
		}(w, r)
	}))
	defer mockHttp.Close()
	url = mockHttp.URL

	r := New(&_http.Client{})
	return r.GetRandNumber(tc.N)
}

// FailEmptyData got "[]" from dependency.
func (tc TestGetRandNumber) FailEmptyData(t *testing.T) (actual int, err error) {
	defer func(backupURL string) { url = backupURL }(url)

	mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
		func(writer _http.ResponseWriter, request *_http.Request) {
			writer.Write([]byte("[]"))
		}(w, r)
	}))
	defer mockHttp.Close()
	url = mockHttp.URL

	r := New(&_http.Client{})
	return r.GetRandNumber(tc.N)
}

// SuccessWithData got "[1,2,3]" from dependency.
func (tc TestGetRandNumber) SuccessWithData(t *testing.T) (actual int, err error) {
	defer func(backupURL string) { url = backupURL }(url)

	mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
		func(writer _http.ResponseWriter, request *_http.Request) {
			writer.Write([]byte("[1,2,3]"))
		}(w, r)
	}))
	defer mockHttp.Close()
	url = mockHttp.URL

	r := New(&_http.Client{})
	return r.GetRandNumber(tc.N)
}
