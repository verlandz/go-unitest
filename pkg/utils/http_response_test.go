package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	tt "github.com/verlandz/go-unitest/pkg/tester"
)

func TestBuildHttpResp(t *testing.T) {
	type in struct {
		data interface{}
		errs []error
	}

	var tcs = []struct {
		name     string
		input    in
		expected HttpResp
	}{
		{
			name: tt.Name{
				Given: "data, errs",
				When:  "when both aren't nil",
				Then:  "return HttpResp with data",
			}.Construct(),
			input: in{
				data: "Hello World",
				errs: []error{
					errors.New("This"),
					errors.New("Is"),
					errors.New("Errors"),
				},
			},
			expected: HttpResp{
				Errors: []string{"This", "Is", "Errors"},
				Data:   "Hello World",
			},
		},
		{
			name: tt.Name{
				Given: "data, errs",
				When:  "when both are nil",
				Then:  "return HttpResp with no data",
			}.Construct(),
			input: in{},
			expected: HttpResp{
				Errors: []string{},
				Data:   nil,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := BuildHttpResp(tc.input.data, tc.input.errs)
			tt.Equal(t, tc.expected, actual)
		})
	}
}

func TestJSON(t *testing.T) {
	type in struct {
		data   interface{}
		errs   []error
		status int
	}

	type exp struct {
		body string
		code int
	}

	var tcs = []struct {
		name     string
		input    in
		expected exp
	}{
		{
			name: tt.Name{
				Given: "data, errs",
				When:  "data is valid",
				Then:  "return body and code",
			}.Construct(),
			input: in{
				data:   "a",
				errs:   []error{errors.New("b")},
				status: http.StatusOK,
			},
			expected: exp{
				body: `{"errors":["b"],"data":"a"}`,
				code: http.StatusOK,
			},
		},
		{
			name: tt.Name{
				Given: "data, errs",
				When:  "data is not valid",
				Then:  "return body and code",
			}.Construct(),
			input: in{
				data:   make(chan int),
				status: http.StatusInternalServerError,
			},
			expected: exp{
				body: "json: unsupported type: chan int\n",
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			BuildHttpResp(tc.input.data, tc.input.errs).JSON(w, tc.input.status)

			tt.Equal(t, tc.expected.body, w.Body.String())
			tt.Equal(t, tc.expected.code, w.Code)
		})
	}
}
