package http

import (
	_http "net/http"
	"net/http/httptest"
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type TestGetRandNumber struct{}

func (TestGetRandNumber) FailClientNil(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "client is nil",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		r := New(nil)
		actual, err = r.GetRandNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestGetRandNumber) FailClientNewRequest(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "client new request is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		defer func(backupURL string) { url = backupURL }(url)
		url = "://"

		r := New(&_http.Client{})
		actual, err = r.GetRandNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestGetRandNumber) FailClientDo(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "client do is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		defer func(backupURL string) { url = backupURL }(url)
		url = ""

		r := New(&_http.Client{})
		actual, err = r.GetRandNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestGetRandNumber) FailClientUnmarshall(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "fail to unmarshall",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		defer func(backupURL string) { url = backupURL }(url)

		mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
			func(writer _http.ResponseWriter, request *_http.Request) {
				writer.Write([]byte(""))
			}(w, r)
		}))
		defer mockHttp.Close()
		url = mockHttp.URL

		r := New(&_http.Client{})
		actual, err = r.GetRandNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestGetRandNumber) FailEmptyData(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "resp data length = 0",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		defer func(backupURL string) { url = backupURL }(url)

		mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
			func(writer _http.ResponseWriter, request *_http.Request) {
				writer.Write([]byte("[]"))
			}(w, r)
		}))
		defer mockHttp.Close()
		url = mockHttp.URL

		r := New(&_http.Client{})
		actual, err = r.GetRandNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestGetRandNumber) Success(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "resp data length > 0",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		defer func(backupURL string) { url = backupURL }(url)

		mockHttp := httptest.NewServer(_http.HandlerFunc(func(w _http.ResponseWriter, r *_http.Request) {
			func(writer _http.ResponseWriter, request *_http.Request) {
				writer.Write([]byte("[1,2,3]"))
			}(w, r)
		}))
		defer mockHttp.Close()
		url = mockHttp.URL

		r := New(&_http.Client{})
		actual, err = r.GetRandNumber(mockN)
		expected := 11

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
