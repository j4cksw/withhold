package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"gopkg.in/mgo.v2"
	"log"
)

type Rate struct {
	Rate float64 `bson:"rate"`
}

func main() {
	fmt.Println("Started")

	router := mux.NewRouter().StrictSlash(true)

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	router.HandleFunc("/withold/amount/{amount}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		vars := mux.Vars(r)
		amount, _ := strconv.ParseFloat(vars["amount"], 64)

		withHold := Rate{}

		collection := session.DB("taxes").C("withold")
		collection.Find(nil).One(&withHold)


		b, _ := json.Marshal(map[string]float64{"value": amount* withHold.Rate})

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(b))
	})

	http.ListenAndServe(":3333", router)

}
