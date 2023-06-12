package handler

import (
	"html/template"
	"net/http"

	"andsp.id/andspdev/config/helpers"
	"andsp.id/andspdev/config/variables"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]any)

	tmpl := template.New("index.html").Funcs(template.FuncMap{
		"site_url": helpers.SiteURL,
	})

	tmpl, err := tmpl.ParseFiles(variables.PathViews + "/home/index.html")

	if err != nil {
		panic(err.Error())
	}

	data["title"] = "AndspID v2.0"

	err = tmpl.Execute(w, data)

	if err != nil {
		panic(err.Error())
	}
}
