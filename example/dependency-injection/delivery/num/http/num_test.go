package http

import (
	"testing"
)

func Test_GetLuckyNumber(t *testing.T) {
	TestGetLuckyNumber{}.FailCalcLuckyNumber(t)
	TestGetLuckyNumber{}.Success(t)
}
