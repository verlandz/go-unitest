package num

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func TestGetNumber(t *testing.T) {
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
			actual := GetNumber()
			tt.Equal(t, tc.expected, actual)
		})
	}
}
