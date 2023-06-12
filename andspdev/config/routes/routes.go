package routes

import (
	"net/http"

	"andsp.id/andspdev/config/variables"
	"andsp.id/andspdev/handler"
	"github.com/gorilla/mux"
)

const rootPath = variables.RootPath
const subRoutes = variables.SubRoutes

func Routes() http.Handler {
	serverMux := mux.NewRouter()

	// Index
	serverMux.HandleFunc("/", handler.IndexHandler)

	// Assets File
	serverMux.HandleFunc("/assets/{path:.*}", handler.AssetsHandler)

	// Login
	serverMux.HandleFunc(subRoutes+"/masuk", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	// Lupa Sandi
	serverMux.HandleFunc(subRoutes+"/lupa-sandi", func(w http.ResponseWriter, r *http.Request) {
		// -----
	})

	// Handle Tampilan 404 Not Found
	serverMux.NotFoundHandler = http.HandlerFunc(handler.Error404Handler)

	return serverMux
}
