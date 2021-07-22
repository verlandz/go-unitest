package http

import (
	"testing"
)

func TestCalc(t *testing.T) {
	Test{}.TestCalcFailWhenDoCalcIsFail(t)
	Test{}.TestCalcSuccess(t)
}
