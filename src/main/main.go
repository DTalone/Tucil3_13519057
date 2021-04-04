package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	// "example.com/handler"
	"example.com/graph"
)

func main() {

	

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/Hello", helloHandler)

	log.Println("starting port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)


	
	// for true {
	// 	fmt.Print("Tampilkan Daftar Nama Simpul ? (y/n)")
	// 	var command string
	// 	fmt.Scanln(&command)
	// 	if command == "y" {
	// 		fmt.Println("Daftar Nama Simpul :")
	// 		graph.PrintListNodes(nodes)
	// 	}
	// 	fmt.Print("Masukkan Nama Posisi AWal : ")
	// 	var start string
	// 	fmt.Scanln(&start)
	// 	fmt.Print("Masukkan Nama Posisi Tujuan : ")
	// 	var goal string
	// 	fmt.Scanln(&goal)
	// 	fmt.Println(graf.Astar(start, goal))
	// }
	// --------------------------------------
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Hello, World!")
	fmt.Println("Selamat datang di aplikasi A*")
	fmt.Print("Masukkan nama  file : ")
	var fileName string
	fmt.Scanln(&fileName)
	graf := graph.ReadFile(fileName)
	nodes := graf.GetNodes()

	log.Printf(r.URL.Path)

	if r.URL.Path != "/"{
		
		http.NotFound(w,r)
		return
	}
	
	data := map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}
	tmpl, err := template.ParseFiles("src/views/index.html")
	log.Println(tmpl)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}


	err = tmpl.Execute(w, nodes)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("anajya ajnajnajibauhb"))
}
