package graph

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// STRUKTUR DATA GRAPH

// Atribut Struktur data Graph
type Graph struct {
	adjacencyMatrix map[string]map[string]float64
	totalNodes      int
	nodes           map[string]Pair
}

// Getter
func (graf *Graph) GetTotalNodes() int {
	return graf.totalNodes
}
func (graf *Graph) GetDistance(A string, B string) float64 {
	return graf.adjacencyMatrix[A][B]
}

// STRUKTUR DATA PAIR

// Atribut Struktur data pair
type Pair struct {
	x float64
	y float64
}

// Getter
func (pair Pair) GetX() float64 {
	return pair.x
}
func (pair Pair) GetY() float64 {
	return pair.y
}

// Error classification
func Check(fileName string, e error) {
	if e != nil {
		log.Fatalf(fileName + ": The system cannot find the file specified.")
	}
}

// Read External File
func ReadFile(fileName string) *Graph {
	graf := new(Graph)
	fileName = "../../test/" + fileName
	file, err := os.Open(fileName)
	Check(fileName, err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// read total nodes
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	graf.totalNodes = n
	// read detail nodes
	graf.nodes = make(map[string]Pair)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		simpang := line[0]
		for j := 1; j < len(line)-2; j++ {
			simpang = simpang + " " + line[j]
		}
		x, _ := strconv.ParseFloat(line[len(line)-2], 64)
		y, _ := strconv.ParseFloat(line[len(line)-1], 64)
		graf.nodes[simpang] = Pair{x, y}
	}

	// Read & Split adjacency matrix
	graf.adjacencyMatrix = make(map[string]map[string]float64)
	for k1 := range graf.nodes {
		graf.adjacencyMatrix[k1] = make(map[string]float64)
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		j := 0
		for k2 := range graf.nodes {
			graf.adjacencyMatrix[k1][k2], _ = strconv.ParseFloat(line[j], 64)
		}
	}
	// for i := 0; i < n; i++ {
	// 	scanner.Scan()
	// 	line := strings.Fields(scanner.Text())
	// 	for j := 0; j < len(line); j++ {
	// 	}
	// }
	return graf
}
