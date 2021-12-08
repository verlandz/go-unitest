package redis

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

func Test_GetTodayNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "client is nil",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetTodayNumber{N: mockN}
		actual, err := tc.FailClientNil(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "get is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetTodayNumber{N: mockN}
		actual, err := tc.FailClientGet(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "the response is bad",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetTodayNumber{N: mockN}
		actual, err := tc.FaiBadResponse(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "all success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetTodayNumber{N: mockN}
		actual, err := tc.Success(t)
		expected := 40

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
