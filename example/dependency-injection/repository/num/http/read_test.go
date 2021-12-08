package http

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

func Test_GetRandNumber(t *testing.T) {
	t.Run(tt.Name{
		Given: "number",
		When:  "client is nil",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.FailClientNil(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "client new request is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.FailClientNewRequest(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "client do is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.FailClientDo(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "fail to unmarshall",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.FailClientUnmarshall(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "resp data length = 0",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.FailEmptyData(t)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	t.Run(tt.Name{
		Given: "number",
		When:  "resp data length > 0",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockN := tdNum.GetNumber3()

		tc := TestGetRandNumber{N: mockN}
		actual, err := tc.SuccessWithData(t)
		expected := 31

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})
}
