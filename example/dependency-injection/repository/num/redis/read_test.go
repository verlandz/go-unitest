package redis

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func Test_GetTodayNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "client is nil",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestGetTodayNumber{}.FailClientNil(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "get is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestGetTodayNumber{}.FailClientGet(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "the response is bad",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestGetTodayNumber{}.FailBadResponse(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "all success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		actual, err := TestGetTodayNumber{}.Success(t)
		expected := 20

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
