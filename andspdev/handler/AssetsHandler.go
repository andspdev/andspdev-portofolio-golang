package handler

import (
	"net/http"

	"andsp.id/andspdev/config/variables"
	"github.com/gorilla/mux"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]

	filepath := variables.PathAssets + "/" + path

	http.ServeFile(w, r, filepath)
}
