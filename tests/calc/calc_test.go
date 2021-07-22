package calc

import (
	"testing"

	tt "github.com/verlandz/go-unitest/pkg/tester"
)

func TestGetN(t *testing.T) {
	var tcs = []struct {
		name     string
		expected int
	}{
		{
			name: tt.Name{
				Given: "nothing",
				When:  "data is present",
				Then:  "return data",
			}.Construct(),
			expected: 10,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetN()
			tt.Equal(t, tc.expected, actual)
		})
	}
}
