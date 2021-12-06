package http

import (
	"testing"
)

func TestGetLuckyNumber(t *testing.T) {
	Test{}.TestGetLuckyNumberFailWhenCalcLuckyNumberIsFail(t)
	Test{}.TestGetLuckyNumberSuccess(t)
}
