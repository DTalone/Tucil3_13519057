package main

import (
	"fmt"

	"example.com/graph"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Selamat datang di aplikasi A*")
	fmt.Print("Masukkan nama  file : ")
	var fileName string
	fmt.Scanln(&fileName)
	graf := new(graph.Graph)
	graf = graph.ReadFile(fileName)
	//fmt.Println(graf.GetTotalNodes())
}
