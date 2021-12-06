package generate

import (
	"testing"
)

func TestGetSquareNumber(t *testing.T) {
	Test{}.TestGetSquareBadRequest(t)
	Test{}.TestGetSquareNumberSuccess(t)
}
