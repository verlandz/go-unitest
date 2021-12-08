package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	rNumHttp "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http"
	rNumHttpMocks "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/http/mocks"
)

type TestCalcLuckyNumber struct {
	N             int
	GetRandNumber rNumHttp.TestGetRandNumber
}

// FailGetRandNumber fail when GetRandNumber.
func (tc TestCalcLuckyNumber) FailGetRandNumber(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumHttpMocks.NewMockClient(gomock.NewController(t))

	mockNumHttp.EXPECT().
		GetRandNumber(tc.GetRandNumber.N).
		Return(tc.GetRandNumber.FailClientNil(t))

	u := New(mockNumHttp)
	return u.CalcLuckyNumber(tc.N)
}

// Success when everything is success.
func (tc TestCalcLuckyNumber) Success(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumHttpMocks.NewMockClient(gomock.NewController(t))

	mockNumHttp.EXPECT().
		GetRandNumber(tc.GetRandNumber.N).
		Return(tc.GetRandNumber.SuccessWithData(t))

	u := New(mockNumHttp)
	return u.CalcLuckyNumber(tc.N)
}
