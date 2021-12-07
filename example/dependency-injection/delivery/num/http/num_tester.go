package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	tt "github.com/verlandz/go-pkg/tester"
	tdNum "github.com/verlandz/go-unitest/example/dependency-injection/test-data/num"
	uNumGenerate "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	uNumGenerateMocks "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component/mocks"
)

type TestGetLuckyNumber struct{}

func (TestGetLuckyNumber) FailCalcLuckyNumber(t *testing.T) {
	mockCalcComponent := uNumGenerateMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "numComponent.CalcLuckyNumber is fail",
		Then:  "return 500",
	}.Construct(), func(t *testing.T) {
		mockCalcComponent.EXPECT().
			CalcLuckyNumber(10).
			Return(uNumGenerate.TestCalcLuckyNumber{}.FailGetRandNumber(t))

		router := httprouter.New()
		request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
		request.Form = make(url.Values)
		request.Form.Add("n", strconv.Itoa(mockN))
		response := httptest.NewRecorder()

		New(router, mockCalcComponent)
		router.ServeHTTP(response, request)

		expected_code := http.StatusInternalServerError
		expected_resp := "client can't be nil"

		tt.Equal(t, expected_code, response.Code)
		tt.Equal(t, expected_resp, response.Body.String())
	})
}

func (TestGetLuckyNumber) Success(t *testing.T) {
	mockCalcComponent := uNumGenerateMocks.NewMockClient(gomock.NewController(t))
	mockN := tdNum.GetNumber()

	t.Run(tt.Name{
		Given: "number",
		When:  "numComponent.CalcLuckyNumber is success",
		Then:  "return 200",
	}.Construct(), func(t *testing.T) {
		mockCalcComponent.EXPECT().
			CalcLuckyNumber(10).
			Return(uNumGenerate.TestCalcLuckyNumber{}.Success(t))

		router := httprouter.New()
		request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
		request.Form = make(url.Values)
		request.Form.Add("n", strconv.Itoa(mockN))
		response := httptest.NewRecorder()

		New(router, mockCalcComponent)
		router.ServeHTTP(response, request)

		expected_code := http.StatusOK
		expected_resp := "121"

		tt.Equal(t, expected_code, response.Code)
		tt.Equal(t, expected_resp, response.Body.String())
	})
}
