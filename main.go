package main

import (
	mtd "SLAE/Methods"
	hlp "SLAE/Helpers"

	"os"
	"fmt"
)


func main() {
	// Initialize SLAE matrices
	var A [][]float64
	var B []float64
	
	hlp.CheckCLI(os.Args, &A, &B)

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
