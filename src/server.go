package main

import (
	"net/http"
	"fmt"
	"withhold"
)

func main() {
	fmt.Println("Started")
	http.ListenAndServe(":3333", withhold.InitRouter())
}
