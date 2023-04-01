package main

import (
	"fmt"
	"math"
)

// Метод итераций для решения СЛАУ
func Iteration(A [][]float64, B, X0 []float64, eps float64) ([]float64, int) {

	// Размер матрицы
	n := len(B)
	// Вектор ответа
	X := make([]float64, n)
	// Нулевые приближения
	xPrev := X0
	// Максимум итераций
	maxIter := int(1e+8)
	// Текущая итерация
	currentIter := 0

	// Проводить вычисления до достижения максимума итераций
	// Либо до достижения необходимой точности
	for currentIter < maxIter {
		// Каждая строка
		for i := 0; i < n; i++ {
			var sum float64

			// Каждый столбец
			for j := 0; j < n; j++ {
				// Если элемент матрицы не лежит на диагонали
				if j != i {
					// Добавить его в сумму
					sum += A[i][j] * X[j]
				}
			}

			// Вычисление новых приближения
			X[i] = (B[i] - sum) / A[i][i]
		}

		// Вычисление точности как суммы
		var diff float64
		// Для каждого из новых приближений
		for i := 0; i < n; i++ {
			diff += math.Abs(X[i] - xPrev[i])
			// Присвоение новых нулевых приближений
			xPrev[i] = X[i]
		}

		// Проверка на совпадение точности
		if diff < eps {
			break
		}

		// Инкремент итераций
		currentIter++
	}

	// Если итераций слишком много - вывести ошибку
	if currentIter >= maxIter {
		fmt.Println("Прошло слишком много итераций")
		return make([]float64, n), currentIter
	}

	// Если всё в порядке - вывести решение
	return X, currentIter
}
