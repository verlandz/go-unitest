package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	rNumRedis "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis"
	rNumRedisMocks "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/redis/mocks"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type TestCalcLuckyNumber struct{}

// FailGetTodayNumber fail when GetTodayNumber.
func (TestCalcLuckyNumber) FailGetTodayNumber(t *testing.T) (actual int, err error) {
	mockNumRedis := rNumRedisMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	mockNumRedis.EXPECT().
		GetTodayNumber(110).
		Return(rNumRedis.TestGetTodayNumber{}.FailClientNil(t))

	u := New(mockNumRedis)
	return u.CalcLuckyNumber(mockN)
}

// Success when everything is success.
func (TestCalcLuckyNumber) Success(t *testing.T) (actual int, err error) {
	mockNumRedis := rNumRedisMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	mockNumRedis.EXPECT().
		GetTodayNumber(110).
		Return(rNumRedis.TestGetTodayNumber{}.Success(t))

	u := New(mockNumRedis)
	return u.CalcLuckyNumber(mockN)
}
