package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	uNumComponent "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component"
	uNumComponentMocks "github.com/verlandz/go-unitest/example/dependency-injection/usecase/num/component/mocks"
)

type TestGetLuckyNumber struct {
	N               int
	CalcLuckyNumber uNumComponent.TestCalcLuckyNumber
}

// FailCalcLuckyNumber fail when CalcLuckyNumber.
func (tc TestGetLuckyNumber) FailCalcLuckyNumber(t *testing.T) *httptest.ResponseRecorder {
	mockCalcComponent := uNumComponentMocks.NewMockClient(gomock.NewController(t))

	mockCalcComponent.EXPECT().
		CalcLuckyNumber(tc.CalcLuckyNumber.N).
		Return(tc.CalcLuckyNumber.FailGetRandNumber(t))

	router := httprouter.New()
	request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
	request.Form = make(url.Values)
	request.Form.Add("n", strconv.Itoa(tc.N))
	response := httptest.NewRecorder()

	New(router, mockCalcComponent)
	router.ServeHTTP(response, request)
	return response
}

// Success when everything is success.
func (tc TestGetLuckyNumber) Success(t *testing.T) *httptest.ResponseRecorder {
	mockCalcComponent := uNumComponentMocks.NewMockClient(gomock.NewController(t))

	mockCalcComponent.EXPECT().
		CalcLuckyNumber(tc.CalcLuckyNumber.N).
		Return(tc.CalcLuckyNumber.Success(t))

	router := httprouter.New()
	request, _ := http.NewRequest(http.MethodGet, "/v1/lucky-number", nil)
	request.Form = make(url.Values)
	request.Form.Add("n", strconv.Itoa(tc.N))
	response := httptest.NewRecorder()

	New(router, mockCalcComponent)
	router.ServeHTTP(response, request)
	return response
}
