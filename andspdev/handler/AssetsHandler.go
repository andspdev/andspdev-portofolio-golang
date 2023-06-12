package handler

import (
	"html/template"
	"net/http"
	"os"

	"andsp.id/andspdev/config/variables"
	"github.com/gorilla/mux"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := vars["path"]

	filepath := variables.PathAssets + "/" + path

	fileInfo, err := os.Stat(filepath)

	// Jika file tidak ada atau folder dibuat forbidden
	if os.IsNotExist(err) || fileInfo.IsDir() {
		w.WriteHeader(http.StatusNotFound)

		tpl, err := template.ParseFiles(variables.PathViews + "/errors/404.html")

		if err != nil {
			panic(err.Error())
		}

		data := map[string]string {
			"request_uri": r.RequestURI,
		}

		err = tpl.Execute(w, data)

		if err != nil {
			panic(err.Error())
		}

		return
	}

	http.ServeFile(w, r, filepath)
}
