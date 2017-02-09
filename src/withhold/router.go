package withhold

import (
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/withold/amount/{amount}", handler)

	return router
}
