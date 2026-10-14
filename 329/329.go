package main

/*
Учитывая матрицу целых чисел m x n,
верните длину самого длинного возрастающего пути в матрице.
Из каждой ячейки вы можете двигаться в четырех направлениях:
влево, вправо, вверх или вниз.
Вы не можете двигаться по диагонали или выходить за границу
(т. е. перенос не допускается).

Пример 1:
https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg
Вход: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Выход: 4
Пояснение: Самый длинный возрастающий путь — это [1, 2, 6, 9].

Пример 2:
https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg
Вход: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Выход: 4
Пояснение: Самый длинный возрастающий путь — это [3, 4, 5, 6].
Перемещение по диагонали запрещено.

Пример 3:
Вход: matrix = [[1]]
Выход: 1

Ограничения:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
0 <= matrix[i][j] <= 2^31 - 1
*/

// longestIncreasingPath ищет максимальную длину пути через DFS с кэшем.
// Строгое возрастание исключает циклы; ответ клетки не зависит от пути входа.
// Время O(V), V=rows*cols: каждая клетка вычисляется раз с четырьмя соседями.
// Рабочая память O(V) для dp и стека; результат int O(1).
// matrix не изменяется; по условию есть хотя бы одна строка и колонка.
func longestIncreasingPath(matrix [][]int) int {
    rows, cols := len(matrix), len(matrix[0])
    dp := make([][]int, rows)
    for i := range dp {
        dp[i] = make([]int, cols)
    }

    var dfs func(r, c, prevVal int) int
    // dfs допускает вход в клетку только после меньшего prevVal.
    // Совокупное время O(V); рабочая память O(V), результат O(1).
    // Заполняет общую dp; matrix не изменяется.
    dfs = func(r, c, prevVal int) int {
        // Проверки границ идут раньше чтения matrix[r][c] благодаря ||.
        // Проверка prevVal идёт ДО кэша: запрещённый вход не должен его читать.
        if r < 0 || r >= rows || c < 0 || c >= cols ||
           matrix[r][c] <= prevVal {
            return 0
        }
        // Допустимый путь содержит эту клетку: длина минимум 1.
        // Поэтому 0 действительно обозначает ещё не вычисленный ответ.
        if dp[r][c] != 0 {
            return dp[r][c]
        }

        res := 1
        res = max(res, 1 + dfs(r+1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r-1, c, matrix[r][c]))
        res = max(res, 1 + dfs(r, c+1, matrix[r][c]))
        res = max(res, 1 + dfs(r, c-1, matrix[r][c]))
        dp[r][c] = res
        return res
    }

    // Каждый старт разрешён: все значения неотрицательны, а prevVal=-1.
    maxPath := 0
    for r := range rows {
        for c := range cols {
			if path := dfs(r, c, -1); path > maxPath {
				maxPath = path
			}
        }
    }
    return maxPath
}
