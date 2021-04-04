package handler

import (
	"html/template"
	"log"
	"net/http"
	// "example.com/graph"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/Hello", helloHandler)

	log.Println("starting port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {

		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("../views/index.html")
	log.Println(tmpl)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("anajya ajnajnajibauhb"))
}
