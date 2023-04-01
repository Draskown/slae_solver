package main

import "fmt"

func main() {
	// Описание входной системы уравнений
	A := [][]float64{{2, -4, -1}, {7, -5, 1}, {8, 7, 4}}
	B := []float64{0.04, -7, 15}

	// Задание метода
	method := "S" // "G" | "I" | "S"

	// Вектор решения
	var x []float64

	if method == "G" {
		// Решение методом Гаусса
		x = Gauss(A, B)

		// Вывод решения
		fmt.Printf("Решение: %v\r\n", x)
	} else {
		// Нулевые приближения
		X0 := []float64{0.0, 0.0, 0.0}
		// Точность
		eps := 0.0001
		// Начальное значение итераций
		i := 0

		if method == "I" {
			// Solve the system using Iterations' method
			x, i = Iteration(A, B, X0, eps)
		} else {
			// Решение методом Зейделя
			x, i = Seidel(A, B, X0, eps)
		}

		// Вывод решения
		if i > 0 {
			fmt.Printf("Решение: %v\r\n Итерации: %d", x, i)
		}
	}
}
