package component

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func Test_CalcLuckyNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "GetTodayNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestCalcLuckyNumber{}.FailGetTodayNumber(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "GetTodayNumber is success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestCalcLuckyNumber{}.Success(t)
		expected := 130

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
