package component

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

func Test_CalcLuckyNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "GetRandNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber2()

		tc := TestCalcLuckyNumber{
			N: mockN,
			GetRandNumber: rNumHttp.TestGetRandNumber{
				N: 210,
			},
		}
		actual, err := tc.FailGetRandNumber(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "GetRandNumber is success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber2()

		tc := TestCalcLuckyNumber{
			N: mockN,
			GetRandNumber: rNumHttp.TestGetRandNumber{
				N: 210,
			},
		}
		actual, err := tc.Success(t)
		expected := 421

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
