package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	rNumRedis "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis"
	rNumRedisMocks "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis/mocks"
)

type TestCalcLuckyNumber struct {
	N              int
	GetTodayNumber rNumRedis.TestGetTodayNumber
}

// FailGetTodayNumber fail when GetTodayNumber.
func (tc TestCalcLuckyNumber) FailGetTodayNumber(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumRedisMocks.NewMockClient(gomock.NewController(t))

	mockNumHttp.EXPECT().
		GetTodayNumber(tc.GetTodayNumber.N).
		Return(tc.GetTodayNumber.FailClientNil(t))

	u := New(mockNumHttp)
	return u.CalcLuckyNumber(tc.N)
}

// Success when everything is success.
func (tc TestCalcLuckyNumber) Success(t *testing.T) (actual int, err error) {
	mockNumHttp := rNumRedisMocks.NewMockClient(gomock.NewController(t))

	mockNumHttp.EXPECT().
		GetTodayNumber(tc.GetTodayNumber.N).
		Return(tc.GetTodayNumber.Success(t))

	u := New(mockNumHttp)
	return u.CalcLuckyNumber(tc.N)
}
