package methods

// Gaussian method for SLAE solving
func Gauss(A [][]float64, B []float64) []float64 {
	// Matrix size
	n := len(A)

	// Straight approach
	// Every row w/o the last one
	for k := 0; k < n-1; k++ {
		// Every column of row + 1 to n
		for i := k + 1; i < n; i++ {
			// Calculate the multiplier
			m := A[i][k] / A[k][k]
			// Add the multiplier to the
			// Lower rows of the matrix
			for j := k + 1; j < n; j++ {
				A[i][j] -= m * A[k][j]
			}
			// Add the multiplier to the right side of the matrix
			B[i] -= m * B[k]
			// Set first elements of the rows to zero
			A[i][k] = 0
		}
	}

	// Reverse approach
	X := make([]float64, n)
	// Every row
	for i := n - 1; i >= 0; i-- {
		/* The sum of the row elements
		must be equal to the corresponding
		element of the right side */
		sum := B[i]

		/*The lower value of the 
		step matrix is being used
		to calculate new multipliers*/
		for j := i + 1; j < n; j++ {
			sum -= A[i][j] * X[j]
		}
		X[i] = sum / A[i][i]
	}

	// Return the result
	return X
}
