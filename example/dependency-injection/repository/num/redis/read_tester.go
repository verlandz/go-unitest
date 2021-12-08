package redis

import (
	"errors"
	"testing"

	"github.com/go-redis/redis"
	"github.com/verlandz/go-unitest/example/dependency-injection/consts"
	pkgRedis "github.com/verlandz/go-unitest/example/dependency-injection/pkg/redis"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type TestGetTodayNumber struct{}

// FailClientNil because client is nil.
func (TestGetTodayNumber) FailClientNil(t *testing.T) (actual int, err error) {
	mockN := tdNum.GetNumber()
	r := New(nil)
	return r.GetTodayNumber(mockN)
}

// FailClientGet because fail to Get.
func (TestGetTodayNumber) FailClientGet(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockN := tdNum.GetNumber()

	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("", errors.New("err"))).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(mockN)
}

// FailBadResponse because the response is bad.
func (TestGetTodayNumber) FailBadResponse(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockN := tdNum.GetNumber()

	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("abc", nil)).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(mockN)
}

// Success when everything is success.
func (TestGetTodayNumber) Success(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockN := tdNum.GetNumber()

	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("10", nil)).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(mockN)
}
