package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {

	serverMux := mux.NewRouter()

	// Index
	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// Login
	serverMux.HandleFunc(SubRoutes+"/masuk", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	// Lupa Sandi
	serverMux.HandleFunc(SubRoutes+"/lupa-sandi", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	return serverMux
}
