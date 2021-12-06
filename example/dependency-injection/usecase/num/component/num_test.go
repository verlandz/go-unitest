package component

import (
	"testing"
)

func TestCalcLuckyNumber(t *testing.T) {
	Test{}.TestCalcLuckyNumberFailWhenGetNumberIsFail(t)
	Test{}.TestCalcLuckyNumberSuccess(t)
}
