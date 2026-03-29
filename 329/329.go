package main

/*
Учитывая матрицу целых чисел m x n, верните длину самого длинного возрастающего пути в матрице.
Из каждой ячейки вы можете двигаться в четырех направлениях: влево, вправо, вверх или вниз.
Вы не можете двигаться по диагонали или выходить за границу (т. е. перенос не допускается).

Пример 1:
https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Пояснение: Самый длинный возрастающий путь — это [1, 2, 6, 9].

Пример 2:
https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg
Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Пояснение: Самый длинный возрастающий путь — это [3, 4, 5, 6]. Перемещение по диагонали запрещено.

Пример 3:
Input: matrix = [[1]]
Output: 1

Ограничения:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
0 <= matrix[i][j] <= 2^31 - 1

*/

func longestIncreasingPath(matrix [][]int) int {
    rows, cols := len(matrix), len(matrix[0])
    dp := make([][]int, rows)
    for i := range dp {
        dp[i] = make([]int, cols)
    }

    var dfs func(r, c, prevVal int) int
    dfs = func(r, c, prevVal int) int {
        if r < 0 || r >= rows || c < 0 || c >= cols ||
           matrix[r][c] <= prevVal {
            return 0
        }
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
