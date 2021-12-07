package component

import (
	"testing"
)

func Test_CalcLuckyNumber(t *testing.T) {
	TestCalcLuckyNumber{}.FailGetRandNumber(t)
	TestCalcLuckyNumber{}.Success(t)
}
