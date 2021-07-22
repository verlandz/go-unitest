package tester

import (
	"errors"
	"testing"
)

func TestLazyAssert(t *testing.T) {
	Error(t, errors.New("err"))
	NoError(t, nil)
	Equal(t, "a", "a")
	NotEqual(t, "a", "b")
	Len(t, []string{"a", "b"}, 2)
	Panics(t, func() { panic("a") })
}
