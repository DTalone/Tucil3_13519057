package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
<<<<<<< HEAD
=======

>>>>>>> 07bd0deee0d6703c1aefdf8410f0fa5651999d7f
	"example.com/graph"
)

// type Info struct {
// 	Latitude  float64
// 	Longitude float64
// 	Name      string
// }

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
<<<<<<< HEAD
	

	fmt.Println(nodes[1].GetLatitude())
	// contoh := {"nama": "a1", "nama1": "b2", "nama2": "c3"}
=======
	// newObject := make([]Info, len(nodes))
	// i := 0
	// for _, v := range nodes {
	// 	newObject[i].latitude = v.GetLatitude()
	// 	newObject[i].longitude = v.GetLongitude()
	// 	newObject[i].name = v.GetName()
	// 	i++
	// }
	// contohs := [len(nodes)][3]int

	// for i := 0 ; i < len(contohs); i++{
	// 	contohs[i][0] = nodes[i].latitude
	// 	contohs[i][1] = nodes[i].longitude
	// 	contohs[i][2] = nodes[i].name
	// }

	// fmt.Println(nodes)
	//contoh := {"nama": "a1", "nama1": "b2", "nama2": "c3"}
>>>>>>> 07bd0deee0d6703c1aefdf8410f0fa5651999d7f

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

<<<<<<< HEAD
	err = tmpl.Execute(w, len(nodes))
=======
	err = tmpl.Execute(w, nodes)
>>>>>>> 07bd0deee0d6703c1aefdf8410f0fa5651999d7f
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("anajya ajnajnajibauhb"))
}
