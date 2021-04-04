package handler

import (
	"log"
	"net/http"
	"path"
	"html/template"

	// "example.com/graph"
)

func homeHandler(w http.ResponseWriter, r *http.Request){

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil{
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}


	err = tmpl.Execute(w, nil)
	if err != nil{
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}