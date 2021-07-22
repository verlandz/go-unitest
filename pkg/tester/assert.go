// NOTE: Please add more if needed
package tester

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getPath() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("From %v:%v", file, line)
}

// Error assert err.
func Error(t *testing.T, err error) {
	assert.Error(t, err, getPath())
}

// NoError assert no err.
func NoError(t *testing.T, err error) {
	assert.NoError(t, err, getPath())
}

// Equal assert equality between two items.
func Equal(t *testing.T, expected interface{}, actual interface{}) {
	assert.Equal(t, expected, actual, getPath())
}

// NotEqual assert inequality between two items.
func NotEqual(t *testing.T, expected interface{}, actual interface{}) {
	assert.NotEqual(t, expected, actual, getPath)
}

// Len assert length of object.
func Len(t *testing.T, obj interface{}, len int) {
	assert.Len(t, obj, len, getPath())
}

// Panics assert if function trigger panic.
func Panics(t *testing.T, fn assert.PanicTestFunc) {
	assert.Panics(t, fn, getPath())
}
