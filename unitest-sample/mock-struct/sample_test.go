package mock_struct

import (
	"testing"

	"github.com/golang/mock/gomock"
	tt "github.com/verlandz/go-unitest/pkg/tester"
	mocks "github.com/verlandz/go-unitest/unitest-sample/mock-struct/mocks"
)

func TestDo(t *testing.T) {
	t.Run(tt.Name{
		Given: "mock and name",
		When:  "both data are valid",
		Then:  "return get name",
	}.Construct(), func(t *testing.T) {
		mockPerson := mocks.NewMockPerson(gomock.NewController(t))
		mockName := "John Doe"

		mockPerson.EXPECT().SetName(mockName).Return()
		mockPerson.EXPECT().GetName().Return(mockName)

		actual := Do(mockPerson, mockName)
		expected := mockName

		tt.Equal(t, expected, actual)
	})
}
