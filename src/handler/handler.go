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
	http.HandleFunc("/map", indexHandler)
	http.HandleFunc("/Hello", helloHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var filepath = path.Join("views", "index.html")
        var tmpl = template.Must(template.New("result").ParseFiles(filepath))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var fileName = r.FormValue("fileName")
		fmt.Println(fileName)
		graf := graph.ReadFile(fileName)
		nodes := graf.GetNodes()


        if err := tmpl.Execute(w, nodes); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }




	// fileName := "tes1.txt"
	// graf := graph.ReadFile(fileName)
	// nodes := graf.GetNodes()
	

	// var filepath = path.Join("views", "index.html")
	// var tmpl, err = template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = tmpl.Execute(w, nodes)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("anajya ajnajnajibauhb"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var filepath = path.Join("views", "home.html")
        var tmpl = template.Must(template.New("form").ParseFiles(filepath))
        var err = tmpl.Execute(w, nil)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }


	var filepath = path.Join("views", "home.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, "", http.StatusBadRequest)
}