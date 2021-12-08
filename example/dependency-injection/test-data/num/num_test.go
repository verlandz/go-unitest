package num

import (
	"testing"

	tt "github.com/verlandz/go-pkg/tester"
)

func TestGetNumber1(t *testing.T) {
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
			actual := GetNumber1()
			tt.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetNumber2(t *testing.T) {
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
			expected: 20,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetNumber2()
			tt.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetNumber3(t *testing.T) {
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
			expected: 30,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := GetNumber3()
			tt.Equal(t, tc.expected, actual)
		})
	}
}
