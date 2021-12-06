package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// GetLuckyNumber gets lucky number with some calculation.
func (d *delivery) GetLuckyNumber(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// request
	n, _ := strconv.Atoi(r.FormValue("n"))

	res, err := d.numComponent.CalcLuckyNumber(n)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, res)
	return
}
