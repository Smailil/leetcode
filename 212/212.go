package main

/*
Дана доска символов board размером m x n и список строк words.
Верните все слова из words, которые есть на доске.
Каждое слово должно состоять из букв последовательно соседних ячеек,
где соседние ячейки соседствуют по горизонтали или по вертикали.
Одна и та же буквенная ячейка не может использоваться в слове более одного раза.

Пример 1:
https://assets.leetcode.com/uploads/2020/11/07/search1.jpg
Вход: board = [["o","a","a","n"],["e","t","a","e"],
["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Выход: ["eat","oath"]

Пример 2:
https://assets.leetcode.com/uploads/2020/11/07/search2.jpg
Вход: board = [["a","b"],["c","d"]], words = ["abcb"]
Выход: []

Ограничения:
m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] — строчная английская буква.
1 <= words.length <= 3 * 10^4
1 <= words[i].length <= 10
words[i] состоит из строчных английских букв.
Все строки words уникальны.
*/

type trieNode struct {
	children map[byte]*trieNode
	isWord bool
}

// newTrieNode создаёт узел с пустой таблицей детей и без отметки слова.
// Время O(1), память нового узла/таблицы O(1).
func newTrieNode() *trieNode {
	return &trieNode{make(map[byte]*trieNode), false}
}

type trie struct {
	root *trieNode
}

// newTrie создаёт независимый trie с одним корневым узлом.
// Время O(1), память O(1); возвращаемая оболочка хранит ссылку на корень.
func newTrie() trie {
	return trie{newTrieNode()}
}

// addWord добавляет путь букв и отмечает конец слова.
// k=len(word): среднее время O(k), новая память trie не больше O(k).
// Временная память O(1); меняет дерево через t.root, не строку word.
func (t trie) addWord(word string) {
	cur := t.root
	// По условию буквы ASCII: rune c точно представима byte(c).
	// Общий префикс использует прежние узлы, недостающие создаём один раз.
	for _, c := range word {
		bc := byte(c)
		if _, ok := cur.children[bc]; !ok {
			cur.children[bc] = newTrieNode()
		}
		cur = cur.children[bc]
	}
	cur.isWord = true
}

type coords struct {
	row int
	col int
}

// newCoords возвращает пару координат: время и память результата O(1).
// Внешнее состояние не меняется; пара служит сравнимым ключом map.
func newCoords(row, col int) coords {
	return coords{row, col}
}

// findWords ищет словарные пути DFS, отсекая отсутствующие префиксы trie.
// R=len(board), C=len(board[0]), B=R*C; L — максимум длины слова,
// S — сумма длин words, D=min(L,B). T — число вызовов DFS:
// среднее время O(S+T*D), включая склейку и хеширование строк префикса.
// Без эффективных отсечений T=O(B*3^D): после первого шага назад идти нельзя.
// P — число узлов trie, Q — число ответов, F — сумма длин найденных слов:
// рабочая память O(P+Q+D^2), результат O(Q+F). P не больше S+1.
// В O(D^2) входят строки всех префиксов активного пути, не только стек.
// Строки res/out используют те же данные слов: их буферы не считаем дважды.
// Исходные board/words не меняются; порядок ответа зависит от обхода map.
func findWords(board [][]byte, words []string) []string {
    t := newTrie()
	for _, w := range words {
		t.addWord(w)
	}
	rows, cols := len(board), len(board[0])
	res := make(map[string]struct{})
	visit := make(map[coords]struct{})

	var dfs func(int, int, *trieNode, string)
	// dfs продолжает word от trie-узла n через ещё не занятую ячейку (r,c).
	// Для T' вызовов: среднее время O(T'*D), стек/visit O(D).
	// Удержанные строки пути занимают O(D^2), общие ответы учтены выше.
	// Меняет visit/res; board и уже построенный trie не меняет.
	dfs = func(r, c int, n *trieNode, word string) {
		// Из допустимой ячейки шаг на 1 достигает лишь -1 или rows/cols.
		// Отсекаем границы до чтения board; строки прямоугольны по условию.
		if r < 0 || c < 0 || r == rows || c == cols {
			return
		}
		coords := newCoords(r, c)
		if _, okVisit := visit[coords]; okVisit {
			return 
		}
		// Если буква не продолжает префикс ни одного слова, вся ветвь лишняя.
		newN, okChildren := n.children[board[r][c]]
		if !okChildren {
			return
		}
		// visit — только текущий путь: одна ячейка не используется в нём дважды.
		// Длина word совпадает с глубиной trie и растёт не выше D.
		visit[coords] = struct{}{}
		n = newN
		word += string(board[r][c])
		// Конец слова может быть также префиксом длинного слова: поиск продолжаем.
		// map убирает повторы одного слова, найденного разными путями.
		if n.isWord {
			res[word] = struct{}{}
		}
		dfs(r-1, c, n, word)
		dfs(r+1, c, n, word)
		dfs(r, c-1, n, word)
		dfs(r, c+1, n, word)
		// Откат освобождает ячейку для других путей, не для текущего же слова.
		delete(visit, coords)
	}
	for r := range rows {
		for c := range cols {
			dfs(r, c, t.root, "")
		}
	}
	// Копируем только строковые заголовки: данные найденных слов уже в res.
	// После возврата out удерживает эти данные, сама map больше не нужна.
	out := make([]string, 0, len(res))
	for w := range res {
		out = append(out, w)
	}
	return out
}
