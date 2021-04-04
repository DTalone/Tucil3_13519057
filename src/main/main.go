package main

import (
	"fmt"

	"example.com/graph"
	"example.com/handler"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Selamat datang di aplikasi A*")
	fmt.Print("Masukkan nama  file : ")
	var fileName string
	fmt.Scanln(&fileName)
	graf := graph.ReadFile(fileName)
	nodes := graf.GetNodes()
	// for true {
	fmt.Print("Tampilkan Daftar Nama Simpul ? (y/n)")
	var command string
	fmt.Scanln(&command)
	if command == "y" {
		fmt.Println("Daftar Nama Simpul :")
		graph.PrintListNodes(nodes)
	}
	fmt.Print("Masukkan Nama Posisi AWal : ")
	var start string
	fmt.Scanln(&start)
	fmt.Print("Masukkan Nama Posisi Tujuan : ")
	var goal string
	fmt.Scanln(&goal)
	fmt.Println(graf.Astar(start, goal))
	// }
	//--------------------------------------
	handler.Start()
}
