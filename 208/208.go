package main

/*
Trie (произносится как «try») или префиксное дерево — это древовидная структура данных, используемая для эффективного хранения и извлечения ключей в наборе строк.
Существуют различные приложения этой структуры данных, такие как автозаполнение и проверка орфографии.

Реализуйте класс Trie:
Trie() Инициализирует объект дерева.
void Insert(String word) Вставляет строковое слово в дерево.
boolean search(String word) Возвращает true, если строковое слово находится в дереве (т. е. было вставлено ранее), и false в противном случае.
логическое значение startWith(String prefix) Возвращает true, если ранее было вставлено строковое слово с префиксом-префиксом, и false в противном случае.

Пример 1:
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]
Объяснение
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

Ограничения:
1 <= word.length, prefix.length <= 2000
word и prefix состоят только из строчных английских букв.
Всего будет выполнено не более 3 *10^4 вызовов для вставки, поиска и запуска.

*/

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

func TrieNodeConstructor() TrieNode {
	return TrieNode{children: make(map[rune]*TrieNode), endOfWord: false}
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	root := TrieNodeConstructor()
	return Trie{root: &root}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		if _, ok := cur.children[c]; ok != true {
			newNode := TrieNodeConstructor()
			cur.children[c] = &newNode
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if _, ok := cur.children[c]; ok != true {
			return false
		}
		cur = cur.children[c]
	}
	return cur.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
		cur := this.root
	for _, c := range prefix {
		if _, ok := cur.children[c]; ok != true {
			return false
		}
		cur = cur.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
