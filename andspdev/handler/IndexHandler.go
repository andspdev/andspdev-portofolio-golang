package handler

import (
	"html/template"
	"net/http"

	"andsp.id/andspdev/config/helpers"
	"andsp.id/andspdev/config/variables"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string] any
	data = make(map[string] any)

	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"site_url": helpers.SiteURL,
	})

	tmpl, err := tmpl.ParseFiles(variables.RootPath + "/views/home/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	data["title"] = "AndspID v2.0"

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
