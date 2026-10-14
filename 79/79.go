package main

/*
Дана сетка символов m x n и строковое слово,
верните true, если слово существует в сетке.
Слово может быть составлено из букв последовательно соседних ячеек,
где соседние ячейки являются соседними по горизонтали или по вертикали.
Одна и та же буквенная ячейка не может использоваться более одного раза.

Пример 1:
https://assets.leetcode.com/uploads/2020/11/04/word2.jpg
Вход: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],
word = "ABCCED"
Выход: true

Пример 2:
https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg
Вход: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],
word = "SEE"
Выход: true

Пример 3:
https://assets.leetcode.com/uploads/2020/10/15/word3.jpg
Вход: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],
word = "ABCB"
Выход: false

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

// exist перебирает старт и ищет word DFS с откатом занятых координат.
// m = len(board), n = len(board[0]), L = len(word).
// Время O(m*n*3^L) в среднем: после первого шага нельзя вернуться назад;
// поиск пары int в path в среднем O(1), без гарантии худшего случая map.
// Рабочая память O(min(L,m*n)) для path и рекурсии, результат — bool.
// board не изменяется; контракт гарантирует непустые board и word.
func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
	path := make(map[coords]struct{})
	var dfs func(int, int, int) bool
	// dfs должен сопоставить word[i:] начиная с (r,c); path содержит
	// ровно клетки уже сопоставленного префикса текущей ветки.
	// q = len(word)-i: время O(3^q) в среднем для оставшегося поиска.
	// Стек O(min(q,rows*cols)+1), общая map path O(min(L,rows*cols)).
	// Меняет только path с откатом, board остаётся прежней.
	dfs = func(r, c, i int) bool {
		// Все буквы уже сопоставлены: следующую координату не читаем,
		// даже если последний вызов шагнул за границу после последней буквы.
		if i == len(word) {
			return true
		}
		// Границы проверяются до обращения board[r][c] и word[i].
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
		// Пометка запрещает повтор той же клетки лишь на этом пути.
		path[coord] = struct{}{}
		res := (dfs(r + 1, c, i+1)) ||
		(dfs(r - 1, c, i+1)) ||
		(dfs(r, c + 1, i+1)) ||
		(dfs(r, c - 1, i+1))
		// После неуспеха откат нужен для следующих путей.
		// Код снимает отметку и при успехе, восстанавливая path,
		// хотя true сразу завершит поиск без другой ветки или старта.
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
