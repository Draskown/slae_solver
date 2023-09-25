package main

import (
	mtd "SLAE/Methods"

	"fmt"
	"os"
	"strconv"
)

// Generate matrices for the SLAE based on the number passed
func generateSLAE(N int) ([][]float64, []float64) {	
	A := make([][]float64, N+1, N+1)
	B := make([]float64, N+1)

	switch N {
	case 1:
		A = [][]float64{
			{3, -0.1},
			{0.1, 7}}
		B = []float64{7.85, -19.3}
	case 2:
		A = [][]float64{
			{1, 1, -1},
			{2, 3, -1},
			{1, 2, 1}}
		B = []float64{-2, 7, 3}
	case 3:
		A = [][]float64{
				{2, -4, -1, 0},
				{7, -5, 1, 3},
				{8, 7, 4, -1},
				{1, 1, 1, 1}}
		B = []float64{0.04, -7, 15, 28}
	case 4:
		A = [][]float64{
				{5, 2, 1, 1, -1},
				{2, -6, 2, 1, 1},
				{1, 2, -5, 1, 1},
				{1, 1, 2, 4, 1},
				{1, -1, 1, 2, 5}}
		B = []float64{-3, 8, 5, -4, 7}
	}

	return A, B
}

func main() {

	// Handle the CLI arguments
	if len(os.Args) > 3 {
		fmt.Println("Too many arguments \nUsage: `go run main.go (1|2|3|4) (G|I|S)`")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Println("Too few arguments \nUsage: `go run main.go (1|2|3|4) (G|I|S)`")
		os.Exit(1)
	}
	
	N, err := strconv.Atoi(os.Args[1])
	if err != nil{
		fmt.Println("Error while reading the Number of the matrices", err)
	}
	
	if N < 1 || N > 4 {
		fmt.Println("Wrong number of the example \nUsage: `go run main.go (1|2|3|4) (G|I|S)`")
		os.Exit(1)
	}
	
	method := os.Args[2]

	if method != "I" && method != "S" && method != "G" {
		fmt.Println("Wrong method \nUsage: `go run main.go (1|2|3|4) (G|I|S)`")
		os.Exit(1)
	}

	A, B := generateSLAE(N)

	// Answer vectror
	var x []float64

	if os.Args[2] == "G" {
		// Solve using gaussian
		x = mtd.Gauss(A, B)

		// Print out the answer
		fmt.Printf("Answer: %v\r\n", x)
	} else {
		// Initial approximations
		X0 := make([]float64, len(A))
		// Accuracy
		eps := 0.001
		// Initial iteration's value
		i := 0

		if os.Args[2] == "I" {
			// Solve the system using iterational method
			x, i = mtd.Iteration(A, B, X0, eps)
		} else {
			// Solve using Seidelian
			x, i = mtd.Seidel(A, B, X0, eps)
		}

		// Print out the answer
		if i > 0 {
			fmt.Printf("Answer: %v\r\nIterations: %d\n", x, i)
		}
	}
}
