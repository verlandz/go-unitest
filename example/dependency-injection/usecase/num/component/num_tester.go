package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	tt "github.com/verlandz/go-pkg/tester"
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
	rNumHttpMocks "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http/mocks"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type TestCalcLuckyNumber struct{}

func (TestCalcLuckyNumber) FailGetRandNumber(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumHttpMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "numHttp.GetRandNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockNumHttp.EXPECT().
			GetRandNumber(110).
			Return(rNumHttp.TestGetRandNumber{}.FailClientNil(t))

		u := New(mockNumHttp)
		actual, err = u.CalcLuckyNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

func (TestCalcLuckyNumber) Success(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumHttpMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "numHttp.GetRandNumber is success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockNumHttp.EXPECT().
			GetRandNumber(110).
			Return(rNumHttp.TestGetRandNumber{}.Success(t))

		u := New(mockNumHttp)
		actual, err = u.CalcLuckyNumber(mockN)
		expected := 121

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
