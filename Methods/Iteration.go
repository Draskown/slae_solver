package methods

import (
	"fmt"
	"math"
)

// Iterational method for SLAE solving
func Iteration(A [][]float64, B, X0 []float64, eps float64) ([]float64, int) {

	// Size of the matrix
	n := len(B)
	// Answer vector
	X := make([]float64, n)
	// Initial approximations
	xPrev := X0
	// Max of iterations
	maxIter := int(1e+8)
	// Current iteration
	currentIter := 0

	// Compute until the number of max iterations are met
	// Or until the needed accuracy is achieved
	for currentIter < maxIter {
		// Every row
		for i := 0; i < n; i++ {
			var sum float64

			// Every column
			for j := 0; j < n; j++ {
				// If an element of the matrix is diagonal
				if j != i {
					// Add it to the sum
					sum += A[i][j] * X[j]
				}
			}

			// Calculate new approximations
			X[i] = (B[i] - sum) / A[i][i]
		}

		// Computing accuracy as a sum
		var diff float64
		// For every of the new approximations
		for i := 0; i < n; i++ {
			diff += math.Abs(X[i] - xPrev[i])
			// Assign new initial approximations
			xPrev[i] = X[i]
		}

		// Check the accuracy
		if diff < eps {
			break
		}

		// Increment current iteration
		currentIter++
	}

	// If there were too many iterations - return an error
	if currentIter >= maxIter {
		fmt.Println("Прошло слишком много итераций")
		return make([]float64, n), currentIter
	}

	// If not - return the answer
	return X, currentIter
}
