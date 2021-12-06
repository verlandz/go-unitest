package generate

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type Test struct{}

// TestGetSquareBadRequest tests func when request is bad.
func (Test) TestGetSquareBadRequest(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNegativeN()

	t.Run(tt.Name{
		Given: "number",
		When:  "number is negative number",
		Then:  "return zero value and err",
	}.Construct(), func(t *testing.T) {
		r := New()
		actual, err = r.GetSquareNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

// TestGetSquareNumberSuccess  tests func when everything is success.
func (Test) TestGetSquareNumberSuccess(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetPositiveN()

	t.Run(tt.Name{
		Given: "number",
		When:  "number is positive number",
		Then:  "return new value and no err",
	}.Construct(), func(t *testing.T) {
		r := New()
		actual, err = r.GetSquareNumber(mockN)
		expected := 100

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
