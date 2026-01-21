package main

/*
Дана сетка символов m x n и строковое слово, верните true, если слово существует в сетке.
Слово может быть составлено из букв последовательно соседних ячеек, где соседние ячейки являются соседними по горизонтали или по вертикали. 
Одна и та же буквенная ячейка не может использоваться более одного раза.

Пример 1:
https://assets.leetcode.com/uploads/2020/11/04/word2.jpg
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true

Пример 2:
https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true

Пример 3:
https://assets.leetcode.com/uploads/2020/10/15/word3.jpg
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false

Ограничения:
m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
доска и слово состоят только из строчных и прописных английских букв.

*/

type coords struct {
	row int
	col int
}

func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
	path := make(map[coords]struct{})
	var dfs func(int, int, int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}
		if (r < 0 || c < 0 || r >= rows || c >= cols) {
			return false
		}
		if (word[i] != board[r][c]) {
			return false
		}
		coord := coords{r, c}
		if _, ok := path[coord]; ok {
			return false
		}
		path[coord] = struct{}{}
		res := (dfs(r + 1, c, i+1)) ||
		(dfs(r - 1, c, i+1)) ||
		(dfs(r, c + 1, i+1)) ||
		(dfs(r, c - 1, i+1))
		delete(path, coord)
		return res
	}
	for r := range rows {
		for c := range cols {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
