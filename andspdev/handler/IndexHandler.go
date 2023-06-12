package handler

import (
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, root string) {

	temp, err := template.ParseFiles(root+"/views/home/index.html")
	
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
	
}