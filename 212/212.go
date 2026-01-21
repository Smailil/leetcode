package main

/*
Дана доска символов m x n и список строк слов, возвращаются все слова на доске.
Каждое слово должно быть составлено из букв последовательно соседних ячеек, где соседние ячейки являются соседними по горизонтали или по вертикали. 
Одна и та же буквенная ячейка не может использоваться в слове более одного раза.

Example 1:
https://assets.leetcode.com/uploads/2020/11/07/search1.jpg
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

Example 2:
https://assets.leetcode.com/uploads/2020/11/07/search2.jpg
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []

Ограничения:
m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] — строчная английская буква.
1 <= words.length <= 3 * 10^4
1 <= words[i].length <= 10
word[i] состоит из строчных английских букв.
Все строки слов уникальны.

*/

type trieNode struct {
	children map[byte]*trieNode
	isWord bool
}

func newTrieNode() *trieNode {
	return &trieNode{make(map[byte]*trieNode), false}
}

type trie struct {
	root *trieNode
}

func newTrie() trie {
	return trie{newTrieNode()}
}

func (t trie) addWord(word string) {
	cur := t.root
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

func newCoords(row, col int) coords {
	return coords{row, col}
}

func findWords(board [][]byte, words []string) []string {
    t := newTrie()
	for _, w := range words {
		t.addWord(w)
	}
	rows, cols := len(board), len(board[0])
	res := make(map[string]struct{})
	visit := make(map[coords]struct{})

	var dfs func(int, int, *trieNode, string)
	dfs = func(r, c int, n *trieNode, word string) {
		if r < 0 || c < 0 || r == rows || c == cols {
			return
		}
		coords := newCoords(r, c)
		if _, okVisit := visit[coords]; okVisit {
			return 
		}
		newN, okChildren := n.children[board[r][c]]
		if !okChildren {
			return
		}
		visit[coords] = struct{}{}
		n = newN
		word += string(board[r][c])
		if n.isWord {
			res[word] = struct{}{}
		}
		dfs(r-1, c, n, word)
		dfs(r+1, c, n, word)
		dfs(r, c-1, n, word)
		dfs(r, c+1, n, word)
		delete(visit, coords)
	}
	for r := range rows {
		for c := range cols {
			dfs(r, c, t.root, "")
		}
	}
	out := make([]string, 0, len(res))
	for w := range res {
		out = append(out, w)
	}
	return out
}
