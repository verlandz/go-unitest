package component

import (
	"testing"
)

func TestDoCalc(t *testing.T) {
	Test{}.TestDoCalcFailWhenGetNumberIsFail(t)
	Test{}.TestDoCalcSuccess(t)
}
