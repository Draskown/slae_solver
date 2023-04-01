package main

// Метод Гаусса для решения СЛАУ
func Gauss(A [][]float64, B []float64) []float64 {
	// Размер матрицы
	n := len(A)

	// Прямой ход метода Гаусса
	// Каждый ряд без последнего
	for k := 0; k < n-1; k++ {
		// Каждый столбец+1
		for i := k + 1; i < n; i++ {
			// Высчитывается коэффициент
			m := A[i][k] / A[k][k]
			// Прибавление коэффициента к
			// Нижним строчкам матрицы
			for j := k + 1; j < n; j++ {
				A[i][j] -= m * A[k][j]
			}
			// Прибавление коэффициента к правой части матрицы
			B[i] -= m * B[k]
			// Обнуление первых элементов строк
			A[i][k] = 0
		}
	}

	// Обратный ход
	X := make([]float64, n)
	// Каждая строка
	for i := n - 1; i >= 0; i-- {
		/* Сумма элементов строки
		Должна равняться соответствующему
		Элементу правой части */
		sum := B[i]

		/* Подставляется нижнее значение
		Ступенчатой матрицы
		И высчитываются коэффициенты */
		for j := i + 1; j < n; j++ {
			sum -= A[i][j] * X[j]
		}
		X[i] = sum / A[i][i]
	}

	// Возвращается результат
	return X
}