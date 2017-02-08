package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AAAA"))
	})

	http.ListenAndServe(":3333", nil)

}
