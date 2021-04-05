package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"example.com/graph"
)

var graf = graph.ReadFile("tes1.txt")

func Start() {
	http.NewServeMux()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/polyline", polyLineHandler)
	http.HandleFunc("/map", indexHandler)
	http.HandleFunc("/Hello", helloHandler)
	http.HandleFunc("/input", inputHandler)

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


func inputHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		fmt.Println("masuk sini")
		var filepath = path.Join("views", "input.html")
        var tmpl = template.Must(template.New("form").ParseFiles(filepath))
        var err = tmpl.Execute(w, nil)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }


	var filepath = path.Join("views", "input.html")
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


func polyLineHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var filepath = path.Join("views", "polyline.html")
        var tmpl = template.Must(template.New("result").ParseFiles(filepath))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		// graf = graph.ReadFile("tes1.txt")
        var simpulAwal = r.FormValue("simpulAwal")
		var simpulAkhir = r.FormValue("simpulAkhir")
		fmt.Println(simpulAwal)
		fmt.Println("gan2")
		distance, rute := graf.Astar(simpulAwal, simpulAkhir)
		fmt.Println("gan11")
		fmt.Println(distance)
		ruteInfo := graf.GetNodeswithIndex(rute)
		fmt.Println(ruteInfo)
		fmt.Println("gann")
		// fmt.Println(fileName)
		// graf := graph.ReadFile(fileName)
		// nodes := graf.GetNodes()


        if err := tmpl.Execute(w, ruteInfo); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
	fmt.Println("anjay")
}