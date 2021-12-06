package component

import (
	"testing"

	"github.com/golang/mock/gomock"
	tt "github.com/verlandz/go-pkg/tester"
	rNumGenerate "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/generate"
	rNumGenerateMocks "github.com/verlandz/go-unitest/example/dependency-injection/repository/num/generate/mocks"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
)

type Test struct{}

// TestCalcLuckyNumberFailWhenGetNumberIsFail tests func when numGenerate.GetSquareNumber is fail.
func (Test) TestCalcLuckyNumberFailWhenGetNumberIsFail(t *testing.T) (actual int, err error) {
	mockNumGenerator := rNumGenerateMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNegativeN()

	t.Run(tt.Name{
		Given: "number",
		When:  "numGenerate.GetSquareNumber is fail",
		Then:  "return no data and err",
	}.Construct(), func(t *testing.T) {
		mockNumGenerator.EXPECT().
			GetSquareNumber(-90).
			Return(rNumGenerate.Test{}.TestGetSquareBadRequest(t))

		u := New(mockNumGenerator)
		actual, err = u.CalcLuckyNumber(mockN)
		expected := 0

		tt.Equal(t, expected, actual)
		tt.Error(t, err)
	})

	return actual, err
}

// TestCalcLuckyNumberSuccess tests func when everything is success.
func (Test) TestCalcLuckyNumberSuccess(t *testing.T) (actual int, err error) {
	mockNumGenerator := rNumGenerateMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetPositiveN()

	t.Run(tt.Name{
		Given: "number",
		When:  "numGenerate.GetSquareNumber is success",
		Then:  "return data and no err",
	}.Construct(), func(t *testing.T) {
		mockNumGenerator.EXPECT().
			GetSquareNumber(110).
			Return(rNumGenerate.Test{}.TestGetSquareNumberSuccess(t))

		u := New(mockNumGenerator)
		actual, err = u.CalcLuckyNumber(mockN)
		expected := 210

		tt.Equal(t, expected, actual)
		tt.NoError(t, err)
	})

	return actual, err
}
