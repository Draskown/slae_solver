package methods

import (
	"fmt"
	"math"
)

// Seidelian method for SLAE solving
func Seidel(A [][]float64, B, X0 []float64, eps float64) ([]float64, int) {

	// Matrix size
	n := len(A)
	// Answer vector
	X := make([]float64, n)
	// Initial approximations
	xPrev := X0
	// Maximum iterations
	maxIter := int(1e+8)
	// Current iteration
	currentIter := 0

	// Compute until the maximum iterations met
	// Or until the needed currency is achieved
	for currentIter < maxIter {
		// Every row
		for i := 0; i < n; i++ {
			var sum1 float64

			// Every column to the left
			for j := 0; j < i; j++ {
				// Sum1 of every elements left of the diagonal is counted
				// Multiplied by new approximations
				sum1 += A[i][j] * X[j]
			}

			var sum2 float64

			// Every column to the right
			for j := i + 1; j < n; j++ {
				// Sum1 of every elements right of the diagonal is counted
				// Multiplied by new approximations
				sum2 += A[i][j] * xPrev[j]
			}

			// Computing new approximations
			X[i] = (B[i] - sum1 - sum2) / A[i][i]
		}

		// Calculation the accurace as a sum
		diff := 0.0
		// For a new of the new approximations
		for i := 0; i < n; i++ {
			diff += math.Abs(X[i] - xPrev[i])
			// Assign new initial approximations
			xPrev[i] = X[i]
		}

		// Check for achieveng the accuracy
		if diff < eps {
			break
		}

		// Increment the current iteration
		currentIter++
	}

	// If there were to many iterations - return an error
	if currentIter >= maxIter {
		fmt.Println("Прошло слишком много итераций")
		return []float64{0.0, 0.0, 0.0}, currentIter
	}

	// If not - return the answer
	return X, currentIter
}
