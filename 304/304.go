package main

/*
Дана двумерная матрица matrix. Необходимо обрабатывать несколько запросов следующего типа:

Вычислить сумму элементов матрицы внутри прямоугольника,
заданного верхним левым углом (row1, col1) и нижним правым углом (row2, col2).
Реализовать класс NumMatrix:

NumMatrix(int[][] matrix) — инициализирует объект с целочисленной матрицей matrix.
int sumRegion(int row1, int col1, int row2, int col2) — возвращает сумму элементов матрицы
внутри прямоугольника, заданного верхним левым углом (row1, col1) и нижним правым углом (row2, col2).
Необходимо разработать алгоритм, в котором sumRegion работает за O(1) по времени.

Пример 1:
https://assets.leetcode.com/uploads/2021/03/14/sum-grid.jpg

Ввод
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
Вывод
[null, 8, 11, 12]

Пояснение
NumMatrix numMatrix = new NumMatrix([[3, 0, 1, 4, 2], [5, 6, 3, 2, 1], [1, 2, 0, 1, 5], [4, 1, 0, 1, 7], [1, 0, 3, 0, 5]]);
numMatrix.sumRegion(2, 1, 4, 3); // вернуть 8 (сумма красного прямоугольника)
numMatrix.sumRegion(1, 1, 2, 2); // вернуть 11 (сумма зелёного прямоугольника)
numMatrix.sumRegion(1, 2, 2, 4); // вернуть 12 (сумма синего прямоугольника)

Ограничения:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-10^4 <= matrix[i][j] <= 10^4
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
Не более 10^4 вызовов sumRegion.

*/

// NumMatrix хранит матрицу 2D префиксных сумм для быстрого вычисления сумм подматриц.
type NumMatrix struct {
	sumMat [][]int // sumMat[i][j] = сумма элементов исходной матрицы в прямоугольнике от (0,0) до (i-1, j-1)
}

// Constructor инициализирует структуру, вычисляя 2D префиксные суммы.
func Constructor(matrix [][]int) NumMatrix {
	rows, cols := len(matrix), len(matrix[0])
	
	// Создаём матрицу префиксных сумм размером (rows+1) x (cols+1).
	// Дополнительная нулевая строка и столбец упрощают граничные случаи и избавляют от проверок на выход за границы.
	sumMat := make([][]int, rows+1)
	for i := range sumMat {
		sumMat[i] = make([]int, cols+1)
	}

	// Заполняем таблицу префиксных сумм.
	// Оптимизация: вместо стандартной формулы с 4 слагаемыми, используем накопление по строкам + сумму сверху.
	for r := range rows {
		prefix := 0 // Сумма элементов текущей строки от 0 до текущего столбца c
		for c := range cols {
			prefix += matrix[r][c]
			
			// above хранит сумму прямоугольника над текущей ячейкой: sumMat[r][c+1]
			above := sumMat[r][c+1]
			
			// Сумма прямоугольника (0,0) -> (r,c) равна сумме текущей строки до c + сумме всего, что выше.
			sumMat[r+1][c+1] = prefix + above
		}
	}

	return NumMatrix{sumMat: sumMat}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// Смещаем индексы на +1, так как sumMat использует 1-based indexing из-за нулевой строки/столбца.
	row1++; col1++; row2++; col2++
	
	// Применяем принцип включения-исключения:
	bottomRight := this.sumMat[row2][col2]      // Сумма всего прямоугольника от (0,0) до (row2, col2)
	above       := this.sumMat[row1-1][col2]    // Вычитаем область выше целевой подматрицы
	left        := this.sumMat[row2][col1-1]    // Вычитаем область левее целевой подматрицы
	topLeft     := this.sumMat[row1-1][col1-1]  // Прибавляем пересечение, которое было вычтено дважды
	
	return bottomRight - above - left + topLeft
}

/**
 * Объект NumMatrix будет создан и вызван следующим образом:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
