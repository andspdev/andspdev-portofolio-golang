package config

import (
	"net/http"

	"andsp.id/andspdev/handler"
	"github.com/gorilla/mux"
)

func Routes() http.Handler {

	serverMux := mux.NewRouter()

	// Index
	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.IndexHandler(w, r, RootPath)
	})

	// Login
	serverMux.HandleFunc(SubRoutes+"/masuk", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	// Lupa Sandi
	serverMux.HandleFunc(SubRoutes+"/lupa-sandi", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	// Handle Tampilan 404 Not Found
	serverMux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.Error404Handler(w, r, RootPath)
	})

	return serverMux
}
