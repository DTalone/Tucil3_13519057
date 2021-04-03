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
	graf := graph.ReadFile(fileName)
	nodes := graf.GetNodes()
	for k, v := range nodes {
		fmt.Println(k, v)
	}
}
