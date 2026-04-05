package main

/*
Дана целочисленая матрица размера m x n, если элемент равен 0, установите всю ее строку и столбец в 0.
Вы должны сделать это на месте.

Пример 1:
https://assets.leetcode.com/uploads/2020/08/17/mat1.jpg
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Пример 2:
https://assets.leetcode.com/uploads/2020/08/17/mat2.jpg
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Ограничения:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1

*/

func setZeroes(matrix [][]int) {
    ROWS, COLS := len(matrix), len(matrix[0])
    rowZero := false

    for r := range ROWS {
        for c := range COLS {
            if matrix[r][c] == 0 {
                matrix[0][c] = 0
                if r > 0 {
                    matrix[r][0] = 0
                } else {
                    rowZero = true
                }
            }
        }
    }

    for r := 1; r < ROWS; r++ {
        for c := 1; c < COLS; c++ {
            if matrix[0][c] == 0 || matrix[r][0] == 0 {
                matrix[r][c] = 0
            }
        }
    }

    if matrix[0][0] == 0 {
        for r := range ROWS {
            matrix[r][0] = 0
        }
    }

    if rowZero {
        for c := range COLS {
            matrix[0][c] = 0
        }
    }
}
