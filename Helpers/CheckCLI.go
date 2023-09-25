package helpers

import (
	"strconv"
)

func CheckCLI (argv []string, A *[][]float64, B *[]float64) {
	// Handle the CLI arguments
	if len(argv) > 3 {
		panic("Too many arguments\nUsage: `go run main.go (1|2|3|4|C) (G|I|S)`")
	}
	if len(argv) < 3 {
		panic("Too few arguments\nUsage: `go run main.go (1|2|3|4|C) (G|I|S)`")
	}
	
	// If a number has been passed
	if argv[1] != "C" {
		N, err := strconv.Atoi(argv[1])
		if err != nil{
			panic("Error while reading the Number of the matrices" + string(err.Error()))
		}
		if N < 1 || N > 4 {
			panic("Wrong number of the example\nUsage: `go run main.go (1|2|3|4|C) (G|I|S)`")
		}

		*A, *B = ChooseSLAE(N)
	} else {
		// If 'C' has been passed
		*A, *B = GenerateSLAE()
	}

	// Check if the supported method was passed
	method := argv[2]
	if method != "I" && method != "S" && method != "G" {
		panic("Wrong method\nUsage: `go run main.go (1|2|3|4|C) (G|I|S)`")
	}
}