package redis

import (
	"errors"
	"testing"

	"github.com/go-redis/redis"
	"github.com/verlandz/go-unitest/example/dependency-injection/consts"
	pkgRedis "github.com/verlandz/go-unitest/example/dependency-injection/pkg/redis"
)

type TestGetTodayNumber struct {
	N int
}

// FailClientNil because client is nil.
func (tc TestGetTodayNumber) FailClientNil(t *testing.T) (actual int, err error) {
	r := New(nil)
	return r.GetTodayNumber(tc.N)
}

// FailClientGet because fail to Get.
func (tc TestGetTodayNumber) FailClientGet(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("", errors.New("err"))).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(tc.N)
}

// FaiBadResponse because the response is bad.
func (tc TestGetTodayNumber) FaiBadResponse(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("abc", nil)).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(tc.N)
}

// Success when everything is success.
// Get 10 from redis.
func (tc TestGetTodayNumber) Success(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedis.InitMock()
	mockRedis.
		On("Get", consts.REDIS_TODAY_NUMBER_KEY).
		Return(redis.NewStringResult("10", nil)).
		Once()

	r := New(mockRedis)
	return r.GetTodayNumber(tc.N)
}
