package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	fmt.Println("Started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		b, _ := json.Marshal(map[string]int{"value": 30})

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(b))
	})

	http.ListenAndServe(":3333", nil)

}
