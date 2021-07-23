package redis

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-redis/redis"
	pkgRedisMocks "github.com/verlandz/go-unitest/pkg/redis/mocks"
	tt "github.com/verlandz/go-unitest/pkg/tester"
	tCalc "github.com/verlandz/go-unitest/tests/calc"
)

type Test struct{}

// TestGetNumberFailWhenRedisIsNil fail when redis is nil.
func (Test) TestGetNumberFailWhenRedisIsNil(t *testing.T) (actual int, err error) {
	mockN := tCalc.GetN()

	t.Run(tt.Name{
		Given: "number",
		When:  "redis is nil",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		r := New(nil)
		actual, err = r.GetNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

// TestGetNumberFailWhenRedisIsFail when redis.Get is fail.
func (Test) TestGetNumberFailWhenRedisIsFail(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedisMocks.NewMock()
	mockN := tCalc.GetN()
	mockKey := fmt.Sprintf(REDIS_GET_NUMBER_KEY, mockN)

	t.Run(tt.Name{
		Given: "number",
		When:  "redis.Get is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockRedis.
			On("Get", mockKey).
			Return(redis.NewStringResult("", errors.New("err")))

		r := New(mockRedis)
		actual, err = r.GetNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

// TestGetNumberSuccess success when everything is ok.
func (Test) TestGetNumberSuccess(t *testing.T) (actual int, err error) {
	mockRedis := pkgRedisMocks.NewMock()
	mockN := tCalc.GetN()
	mockKey := fmt.Sprintf(REDIS_GET_NUMBER_KEY, mockN)

	t.Run(tt.Name{
		Given: "number",
		When:  "everything is ok",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockRedis.
			On("Get", mockKey).
			Return(redis.NewStringResult("10", nil))

		r := New(mockRedis)
		actual, err = r.GetNumber(mockN)
		expected := 10

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
