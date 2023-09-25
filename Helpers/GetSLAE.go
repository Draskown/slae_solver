package helpers

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

// Generate matrices for the SLAE based on the number passed
func ChooseSLAE(N int) ([][]float64, []float64) {	
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


// Helper to read strings with spaces
func readString(str *string) {
	// Create a buffer to hold the input
	in := bufio.NewReader(os.Stdin)

	// If read string works, return the read line without '\n'
	if line, err := in.ReadString('\n'); err == nil{
		*str = line[:len(line)-1]
	} else {
		// Exit the program if not
		panic(err)
	}

}


func GenerateSLAE () ([][]float64, []float64) {
	// Initaliazing of operated values
	var (
		// String user input
		input string
		// Size of the matrix
		size int
	)

	// Read the size of the matrix
	fmt.Println("Enter the size of your matrix: ")
	readString(&input)

	// Convert the string to int
	size, err := strconv.Atoi(input)
	if err != nil {
		// Prompt the user to enter the size again
		// If there was an error of parsing
		fmt.Printf("Wrong size input.\n\n")
		GenerateSLAE()
	}

	// Init the main matrix and B column slices
	matrix := make([][]float64, size)
	column := make([]float64, size)
	
	// Iterate through rows
	for i := 0; i < size + 1; i++ {
		// B column should be entered the last,
		// So while i != size, prompt the user to enter main rows
		if i < size {
			fmt.Printf("Enter the %d row of the matrix: \n", i+1)
			} else {
				fmt.Println("Enter the B column of the matrix: ")
			}
		// Read the input
		readString(&input)
		temp := strings.Split(input, " ")
			
		// Create a temp row in order
		// To write the matrix correctly
		row := make([]float64, size)

		// If the entered row does not equal
		// Nn size to the entered
		if len(temp) < size || len(temp) > size {
			fmt.Printf("The row has different amount of values (%d instead of expected %d).\n" +
						"Please try again.\n\n", len(temp), size)
			// Rebase current iterator
			i--
			// And ask for the prompt again
			// By skipping through the rest of processing
			continue
		}

		// If input has been accepted,
		for j, v := range temp {
			// Parse string to float
			if a, err := strconv.ParseFloat(v, 64); err == nil {
				row[j] = a
			} else {
				// If not - exit the program
				panic(err)
			}
		}
		
		// Check whether it is time to
		// Fill the B column
		if i < size {
			matrix[i] = row
		} else {
			copy(column, row)
		}
	}

	return matrix, column
}