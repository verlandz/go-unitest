package redis

import (
	"testing"
)

func TestGetNumber(t *testing.T) {
	Test{}.TestGetNumberFailWhenRedisIsNil(t)
	Test{}.TestGetNumberFailWhenRedisIsFail(t)
	Test{}.TestGetNumberSuccess(t)
}
