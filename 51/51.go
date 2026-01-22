package main

import "strings"

/*
Загадка с n ферзями — это задача о том, как разместить n ферзей на шахматной доске размера n x n так, чтобы никакие два ферзя не атаковали друг друга.
Учитывая целое число n, найдите все различные решения головоломки с n ферзями. Вы можете вернуть ответ в любом порядке.
Каждое решение содержит отдельную конфигурацию доски для размещения n ферзей, где «Q» и «.». оба указывают на ферзя и пустое место соответственно.

Пример 1:
https://assets.leetcode.com/uploads/2020/11/13/queens.jpg
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Пояснение: существует два различных решения головоломки с четырьмя ферзями, как показано выше.

Пример 2:
Input: n = 1
Output: [["Q"]]

Ограничения:
1 <= n <= 9

*/

func solveNQueens(n int) [][]string {
    col := make(map[int]struct{})
	posDiag := make(map[int]struct{}) // (r+c)
	negDiag := make(map[int]struct{}) // (r-c)
	var res [][]string
	board := make([][]string, n)
	for i := range n {
		board[i] = make([]string, n)
		for j := range n {
			board[i][j] = "."
		}
	}
	var backtrack func(int)
	backtrack = func(r int) {
		if r == n {
			out := make([]string, n)
			for i, row := range board {
				out[i] = strings.Join(row, "")
			}
			res = append(res, out)
			return
		}
		for c := range n {
			if _, ok := col[c]; ok {
				continue
			}
			if _, ok := posDiag[r+c]; ok {
				continue
			}
			if _, ok := negDiag[r-c]; ok {
				continue
			}
			col[c] = struct{}{}
			posDiag[r+c] = struct{}{}
			negDiag[r-c] = struct{}{}
			board[r][c] = "Q"

			backtrack(r+1)

			delete(col, c)
			delete(posDiag, r+c)
			delete(negDiag, r-c)
			board[r][c] = "."
		}
	}
	backtrack(0)
	return res
}
