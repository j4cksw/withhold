package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"withhold"
)

type Rate struct {
	Rate float64 `bson:"rate"`
}

func main() {
	fmt.Println("Started")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/withold/amount/{amount}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		amount := ParseAmount(r)

		b:= PrepareResponseBody(CalculateWithHolding(amount))

		WriteResponse(w, b)
	})

	http.ListenAndServe(":3333", router)

}
func WriteResponse(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ParseAmount(r *http.Request) float64 {
	vars := mux.Vars(r)
	amount, _ := strconv.ParseFloat(vars["amount"], 64)
	return amount
}

func CalculateWithHolding(amount float64) float64{
	return withhold.NewWithHolding().GetWithHolding(amount)
}

func PrepareResponseBody(calculated float64) []byte {
	b, _ := json.Marshal(map[string]float64{ "value": calculated })
	return b
}

