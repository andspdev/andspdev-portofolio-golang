package handler

import (
	"html/template"
	"net/http"

	"andsp.id/andspdev/config/variables"
)

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	temp, err := template.ParseFiles(variables.PathViews + "/errors/404.html")

	data := map[string]string{
		"request_uri": r.RequestURI,
	}

	if err != nil {
		panic(err.Error())
	}

	temp.Execute(w, data)
}
