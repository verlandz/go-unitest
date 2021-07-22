package http

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	utils "github.com/verlandz/go-unitest/pkg/utils"
)

// Calc is feat to do some calc.
func (d *delivery) Calc(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	n, _ := strconv.Atoi(r.FormValue("n"))

	// add more complex (1)
	n *= 10

	res, err := d.calcComponent.DoCalc(n)
	if err != nil {
		utils.BuildHttpResp(nil, []error{err}).JSON(w, http.StatusInternalServerError)
		return
	}

	// add more complex (2)
	res += 2

	utils.BuildHttpResp(res, nil).JSON(w, http.StatusOK)
	return
}
