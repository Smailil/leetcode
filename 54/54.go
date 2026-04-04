package main

/*
Дана матрица размера m x n, верните все элементы матрицы в спиральном порядке.

Пример 1:

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Пример 2:

Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Ограничения:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

*/

func spiralOrder(matrix [][]int) []int {
	// Определяем границы обхода.
	// left/right отвечают за столбцы, top/bottom за строки.
	// right и bottom используют полуоткрытый интервал (не включаются в диапазон),
	// что упрощает условия в циклах for.
	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)

	res := make([]int, 0, right*bottom)

	// Основной цикл продолжается, пока границы не пересекутся
	for left < right && top < bottom {
		// 1. Движение ВПРАВО по верхней строке
		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ // Верхняя строка обработана, сдвигаем границу вниз

		// 2. Движение ВНИЗ по правому столбцу
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right-- // Правый столбец обработан, сдвигаем границу влево

		// Необходима для матриц с одной строкой или одним столбцом,
		// где после первых двух проходов границы уже могут пересечься.
		if !(left < right && top < bottom) {
			break
		}

		// 3. Движение ВЛЕВО по нижней строке
		for i := right - 1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom-- // Нижняя строка обработана, сдвигаем границу вверх

		// 4. Движение ВВЕРХ по левому столбцу
		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++ // Левый столбец обработан, сдвигаем границу вправо
	}

	return res
}
