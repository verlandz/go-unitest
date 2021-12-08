package http

import (
	"net/http"
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
)

func Test_GetLuckyNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "CalcLuckyNumber is fail",
		Then:  "return 500",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber1()

		tc := TestGetLuckyNumber{
			N: mockN,
			CalcLuckyNumber: uNumComponent.TestCalcLuckyNumber{
				N: 10,
				GetRandNumber: rNumHttp.TestGetRandNumber{
					N: 110,
				},
			},
		}
		actual := tc.FailCalcLuckyNumber(t)
		expected_code := http.StatusInternalServerError
		expected_resp := "client can't be nil"

		tt.Equal(t, expected_code, actual.Code)
		tt.Equal(t, expected_resp, actual.Body.String())
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "CalcLuckyNumber is success",
		Then:  "return 200",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber1()

		tc := TestGetLuckyNumber{
			N: mockN,
			CalcLuckyNumber: uNumComponent.TestCalcLuckyNumber{
				N: 10,
				GetRandNumber: rNumHttp.TestGetRandNumber{
					N: 110,
				},
			},
		}
		actual := tc.Success(t)
		expected_code := http.StatusOK
		expected_resp := "221"

		tt.Equal(t, expected_code, actual.Code)
		tt.Equal(t, expected_resp, actual.Body.String())
	})
}
