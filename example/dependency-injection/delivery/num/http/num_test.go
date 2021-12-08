package http

import (
	"net/http"
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func Test_GetLuckyNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "CalcLuckyNumber is fail",
		Then:  "return 500",
	}.Construct(), func(t *testing.T) {
		actual := TestGetLuckyNumber{}.FailCalcLuckyNumber(t)
		expected_code := http.StatusInternalServerError
		expected_resp := "client is nil"

		tt.Equal(t, expected_code, actual.Code)
		tt.Equal(t, expected_resp, actual.Body.String())
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "CalcLuckyNumber is success",
		Then:  "return 200",
	}.Construct(), func(t *testing.T) {
		actual := TestGetLuckyNumber{}.Success(t)
		expected_code := http.StatusOK
		expected_resp := "130"

		tt.Equal(t, expected_code, actual.Code)
		tt.Equal(t, expected_resp, actual.Body.String())
	})
}
