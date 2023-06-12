package handler

import (
	"html/template"
	"net/http"
)

func Error404Handler(w http.ResponseWriter, r *http.Request, root string) {
	w.WriteHeader(http.StatusNotFound)

	temp, err := template.ParseFiles(root+"/views/errors/404.html")

	data := map[string]string {
		"request_uri": r.RequestURI,
	}

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}
