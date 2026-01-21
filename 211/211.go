package main

/*
Разработайте структуру данных, которая поддерживает добавление новых слов и определение совпадения строки с какой-либо ранее добавленной строкой.
Реализуйте класс WordDictionary:
WordDictionary() Инициализирует объект.
void addWord(word) Добавляет слово в структуру данных, его можно сопоставить позже.
bool search(word) Возвращает true, если в структуре данных есть какая-либо строка, соответствующая слову, или false в противном случае. 
Cлово может содержать точки '.' где точкам можно сопоставить любую букву.
 
Example:
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]
Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True

Ограничения:
1 <= word.length <= 25
Слово в addWord состоит из строчных английских букв.
искомое слово состоит из '.' или строчные английские буквы.
В слове для поисковых запросов будет не более 2 точек.
Для addWord и поиска будет выполнено не более 10^4 вызовов.

*/

type trieNode struct {
	children  map[rune]*trieNode
	endOfWord bool
}

func nodeConstructor() trieNode {
    return trieNode{make(map[rune]*trieNode), false}
}

type WordDictionary struct {
    root *trieNode
}


func Constructor() WordDictionary {
    root := nodeConstructor()
	return WordDictionary{root: &root}
}


func (this *WordDictionary) AddWord(word string)  {
    cur := this.root
	for _, c := range word {
		if _, ok := cur.children[c]; !ok {
			newNode := nodeConstructor()
			cur.children[c] = &newNode
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true 
}


func (this *WordDictionary) Search(word string) bool {
	var dfs func(j int, root *trieNode) bool
	dfs = func(j int, root *trieNode) bool {
		cur := root
		for i := j; i < len(word); i++ {
			c := rune(word[i])
			if c == '.' {
				for _, child := range cur.children {
					if dfs(i+1, child) {
						return true
					}
                }
				return false
			} else {
				if _, ok := cur.children[c]; !ok {
					return false
				}
				cur = cur.children[c]
			}
		}
		return cur.endOfWord
	}
	return dfs(0, this.root)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
