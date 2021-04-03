package graph

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// STRUKTUR DATA GRAPH

// Atribut Struktur data Graph
type Graph struct {
	adjacencyMatrix [][]float64
	totalNodes      int
	nodes           []Info
}

// Getter
func (graf *Graph) GetTotalNodes() int {
	return graf.totalNodes
}
func (graf *Graph) GetDistance(A string, B string) float64 {
	idx1, idx2 := Search(A, B, graf.nodes)
	return graf.adjacencyMatrix[idx1][idx2]
}
func (graf *Graph) GetNodes() []Info {
	return graf.nodes
}

func Search(A string, B string, nodes []Info) (int, int) {
	idx1 := -1
	idx2 := -1
	for i := 0; i < len(nodes); i++ {
		if nodes[i].name == A {
			idx1 = i
		}
		if nodes[i].name == A {
			idx2 = i
		}
	}
	return idx1, idx2
}

// STRUKTUR DATA PAIR

// Atribut Struktur data pair
type Info struct {
	latitude  float64
	longitude float64
	name      string
}

// Getter
func (info Info) GetLatitude() float64 {
	return info.latitude
}
func (info Info) GetLongitude() float64 {
	return info.longitude
}
func (info Info) GetName() string {
	return info.name
}

// Error classification
func Check(fileName string, e error) {
	if e != nil {
		log.Fatalf(fileName + ": The system cannot find the file specified.")
	}
}

// Euclidan Distance
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}
func GetEuclidanDistance(A Info, B Info) float64 {
	// Coordinate A in radians
	latitude1 := degreesToRadians(A.latitude)
	longitude1 := degreesToRadians(A.longitude)

	// Coordinate B in radians
	latitude2 := degreesToRadians(B.latitude)
	longitude2 := degreesToRadians(B.longitude)

	// Haversine Formula
	differencelongitude := longitude2 - longitude1
	differencelatitude := latitude2 - latitude1

	a := math.Pow(math.Sin(differencelatitude/2), 2) + math.Cos(latitude1)*math.Cos(latitude2)*math.Pow(math.Sin(differencelongitude/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// earth radius in KM
	km := c * 6371
	return km
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
	graf.nodes = make([]Info, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		simpang := line[0]
		for j := 1; j < len(line)-2; j++ {
			simpang = simpang + " " + line[j]
		}
		latitude, _ := strconv.ParseFloat(line[len(line)-2], 64)
		longitude, _ := strconv.ParseFloat(line[len(line)-1], 64)
		graf.nodes[i] = Info{latitude, longitude, simpang}
	}

	// Read & Split adjacency matrix
	graf.adjacencyMatrix = make([][]float64, n)
	for i := 0; i < n; i++ {
		graf.adjacencyMatrix[i] = make([]float64, n)
		scanner.Scan()
		line := strings.Fields(scanner.Text())
		for j := 0; j < n; j++ {
			if line[j] == "0" {
				graf.adjacencyMatrix[i][j] = 0
			} else {
				a := graf.nodes[i]
				b := graf.nodes[j]
				graf.adjacencyMatrix[i][j] = GetEuclidanDistance(a, b)
			}
		}
	}
	return graf
}
