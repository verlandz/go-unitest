package component

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	rNumRedis "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

func Test_CalcLuckyNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "GetTodayNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber2()

		tc := TestCalcLuckyNumber{
			N: mockN,
			GetTodayNumber: rNumRedis.TestGetTodayNumber{
				N: 210,
			},
		}
		actual, err := tc.FailGetTodayNumber(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "GetTodayNumber is success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber2()

		tc := TestCalcLuckyNumber{
			N: mockN,
			GetTodayNumber: rNumRedis.TestGetTodayNumber{
				N: 210,
			},
		}
		actual, err := tc.Success(t)
		expected := 430

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
