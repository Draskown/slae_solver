# GO SLAE Solver

Solves systems of liner algebric equations using one of the three methods: Gaussian, Iterative and Siedelian.

## Prerequisites

Go 1.20+

## Usage

```
go run main.go (1|2|3|4|C) (G|I|S)
```

The first argument stands for the example SLAE (1 - 2-dimensional, 2 - 3-dimensional, etc...).

Or you can pass `C` instead and enter your own SLAE:

```
> go run main.go C G
Enter the size of your matrix: 
2
Enter the 1 row of the matrix: 
3 -0.1
Enter the 2 row of the matrix: 
0.1 7
Enter the B column of the matrix: 
7.85 -19.3
Answer: [2.5235602094240837 -2.793193717277487]
```

**User entered matrices must be separated by spaces. Otherwise the input will not parse.**

The second argument stands for the method of solving the chosen SLAE (G - Gaussian, I - Iterative, S - Seidelian).

*Example:*

```
> go run main.go 4 S
Answer: [0.6101565277665252 -1.393274696069548 -1.2508517278979838 -0.5457594862600068 1.4677878953163848]
Iterations: 6
```

## Room for improvement

- Parse the input with commas
- Feel free to contribute :0