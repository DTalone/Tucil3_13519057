package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"example.com/graph"
)

func Start() {
	http.NewServeMux()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/Hello", helloHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "tes1.txt"
	graf := graph.ReadFile(fileName)
	nodes := graf.GetNodes()
	

	fmt.Println(nodes[1].GetLatitude())
	// contoh := {"nama": "a1", "nama1": "b2", "nama2": "c3"}

	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// var data = map[string]interface{}{
	// 	"title": "Learning Golang Web",
	// 	"name":  "Batman",
	// }

	err = tmpl.Execute(w, len(nodes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("anajya ajnajnajibauhb"))
}
