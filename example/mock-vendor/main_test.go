package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	tt "github.com/verlandz/go-pkg/tester"
	personMocks "github.com/verlandz/go-unitest/example/mock-vendor/person/mocks"
)

func TestDo(t *testing.T) {
	t.Run(tt.Name{
		Given: "mock and name",
		When:  "both data are valid",
		Then:  "return get name",
	}.Construct(), func(t *testing.T) {
		mockPerson := personMocks.NewMockClient(gomock.NewController(t))
		mockName := "John Doe"

		mockPerson.EXPECT().SetName(mockName).Return()
		mockPerson.EXPECT().GetName().Return(mockName)

		actual := Do(mockPerson, mockName)
		expected := "John Doe"

		tt.Equal(t, expected, actual)
	})
}
