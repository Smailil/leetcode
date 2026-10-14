package main

/*
Trie (произносится как «try») или префиксное дерево — это древовидная
структура данных, используемая для эффективного хранения и извлечения ключей
в наборе строк.
Существуют различные приложения этой структуры данных, такие как автозаполнение
и проверка орфографии.

Реализуйте класс Trie:
Trie() Инициализирует объект дерева.
void insert(String word) Вставляет строковое слово в дерево.
boolean search(String word) Возвращает true, если строковое слово находится
в дереве (то есть было вставлено ранее), и false в противном случае.
boolean startsWith(String prefix) Возвращает true, если ранее было вставлено
слово с префиксом prefix, и false в противном случае.

Пример 1:
Вход:
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Выход:
[null, null, true, false, true, null, true]
Объяснение:
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // возвращает True
trie.search("app");     // возвращает False
trie.startsWith("app"); // возвращает True
trie.insert("app");
trie.search("app");     // возвращает True

Ограничения:
1 <= word.length, prefix.length <= 2000
word и prefix состоят только из строчных английских букв.
Всего будет выполнено не более 3 *10^4 вызовов insert, search и startsWith.
*/

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

// TrieNodeConstructor создаёт узел с пустыми переходами и без конца слова.
// Время O(1), рабочая память и новый узел O(1); внешние данные не меняются.
func TrieNodeConstructor() TrieNode {
	return TrieNode{children: make(map[rune]*TrieNode), endOfWord: false}
}

type Trie struct {
	root *TrieNode
}

// Constructor создаёт пустое префиксное дерево с отдельным корнем.
// Время O(1), рабочая память и результат O(1); внешних изменений нет.
func Constructor() Trie {
	root := TrieNodeConstructor()
	return Trie{root: &root}
}

// Insert достраивает путь word и отмечает конец целого слова.
// Время в среднем O(l), l = len(word): ASCII-ключ rune имеет фиксированный
// размер. Рабочая память O(1), до O(l) новых узлов остаются в дереве.
// Изменяет this; общие префиксы переиспользуются, word не изменяется.
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		// cur соответствует уже пройденному префиксу. Отсутствующий
		// переход создаём один раз; существующий сохраняет другие слова.
		if _, ok := cur.children[c]; ok != true {
			newNode := TrieNodeConstructor()
			cur.children[c] = &newNode
		}
		cur = cur.children[c]
	}
	// Только последний узел означает вставленное слово: промежуточный
	// префикс может существовать, но не быть самостоятельным словом.
	cur.endOfWord = true
}

// Search проходит путь word и проверяет метку завершённого слова.
// Время в среднем O(l), l = len(word), по map с ключами постоянного размера.
// Рабочая память и булев результат O(1); this и word не изменяются.
func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		if _, ok := cur.children[c]; ok != true {
			return false
		}
		cur = cur.children[c]
	}
	// Наличие всех переходов доказывает только префикс, поэтому отдельно
	// проверяем, завершалось ли здесь когда-либо вставленное слово.
	return cur.endOfWord
}

// StartsWith проверяет существование пути prefix, не требуя конца слова.
// Время в среднем O(l), l = len(prefix), при постоянном размере rune.
// Рабочая память и булев результат O(1); this и prefix не изменяются.
func (this *Trie) StartsWith(prefix string) bool {
		cur := this.root
	for _, c := range prefix {
		if _, ok := cur.children[c]; ok != true {
			return false
		}
		cur = cur.children[c]
	}
	// В отличие от Search, любой достигнутый узел завершает нужный
	// префикс; значение endOfWord здесь не имеет значения.
	return true
}

/**
 * Объект Trie будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
