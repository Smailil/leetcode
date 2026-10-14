package main

/*
Разработайте структуру данных, которая поддерживает добавление новых слов
и определение совпадения строки с какой-либо ранее добавленной строкой.
Реализуйте класс WordDictionary:
WordDictionary() Инициализирует объект.
void addWord(word) Добавляет слово в структуру данных,
его можно сопоставить позже.
bool search(word) Возвращает true, если в структуре данных есть строка,
соответствующая word, или false в противном случае.
word может содержать точки '.', каждой из которых можно сопоставить любую букву.

Пример:
Вход:
["WordDictionary","addWord","addWord","addWord",
"search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Выход:
[null,null,null,null,false,true,true,true]
Объяснение:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // возвращает False
wordDictionary.search("bad"); // возвращает True
wordDictionary.search(".ad"); // возвращает True
wordDictionary.search("b.."); // возвращает True

Ограничения:
1 <= word.length <= 25
word в addWord состоит из строчных английских букв.
Искомое word состоит из '.' или строчных английских букв.
В слове для поисковых запросов будет не более 2 точек.
Для addWord и search будет выполнено не более 10^4 вызовов.
*/

// trieNode хранит рёбра префиксного дерева и признак конца слова.
// Одинаковые префиксы разделяются; endOfWord отличает слово от префикса.
type trieNode struct {
	children  map[rune]*trieNode
	endOfWord bool
}

// nodeConstructor создаёт пустую вершину с доступной для записи картой.
// Время, рабочая память и новый результат O(1).
func nodeConstructor() trieNode {
    return trieNode{make(map[rune]*trieNode), false}
}

type WordDictionary struct {
    root *trieNode
}


// Constructor создаёт словарь с одной корневой вершиной без слов.
// Время, рабочая память и начальный результат O(1).
func Constructor() WordDictionary {
    root := nodeConstructor()
	return WordDictionary{root: &root}
}


// AddWord добавляет путь букв и отмечает конечную вершину слова.
// Среднее время O(m), m=len(word), при O(1) операциях map с rune-ключами.
// Рабочая память O(1); новых вершин и рёбер не больше O(m).
// Меняет дерево this; весь словарь занимает O(S), S — число его вершин.
// По условию буквы ASCII, поэтому rune соответствует отдельным буквам.
func (this *WordDictionary) AddWord(word string)  {
    cur := this.root
	for _, c := range word {
		// Переиспользуем общий префикс, создавая только недостающее ребро.
		if _, ok := cur.children[c]; !ok {
			newNode := nodeConstructor()
			cur.children[c] = &newNode
		}
		cur = cur.children[c]
	}
	// Полное слово отмечается и при повторном добавлении того же пути.
	cur.endOfWord = true 
}


// Search ищет полное слово по шаблону; "." заменяет ровно одну букву.
// Среднее время O(m*A^d), m=len(word), d — число точек, A — размер алфавита.
// Здесь A=26, d до 2; операции rune-map средние O(1).
// Рабочая память O(d+1) для рекурсии по точкам; результат bool O(1).
// Дерево не меняется; byte-индексы верны для букв ASCII из условия.
func (this *WordDictionary) Search(word string) bool {
	var dfs func(j int, root *trieNode) bool
	// dfs проверяет хвост от j в переданном поддереве.
	// Среднее время O(m*A^d), рабочая память O(d+1), результат O(1).
	dfs = func(j int, root *trieNode) bool {
		cur := root
		for i := j; i < len(word); i++ {
			c := rune(word[i])
			// Пробуем каждое ребро, потребляя один символ шаблона.
			// Первое полное совпадение завершает перебор остальных ветвей.
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
		// Совпадение префикса недостаточно: здесь должно завершаться слово.
		return cur.endOfWord
	}
	return dfs(0, this.root)
}


/**
 * Объект WordDictionary будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
