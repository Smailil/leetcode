package main

/*
Спроектируйте структуру данных для кэша LRU:
при вытеснении удаляется ключ, который дольше всего не использовался.

Реализуйте класс LRUCache:
LRUCache(int capacity) инициализирует кэш LRU с положительной ёмкостью.
int get(int key) возвращает значение ключа, если он существует,
иначе возвращает -1.
void put(int key, int value) обновляет значение существующего ключа.
В противном случае добавьте пару ключ-значение в кэш.
Если после операции число ключей превышает ёмкость,
удалите ключ, который дольше всего не использовался.
Каждая функция get и put должна работать за O(1) в среднем.

Пример 1:
Вход:
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Выход:
[null, null, null, 1, null, -1, null, -1, 3, 4]
Пояснение:
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); //кэш равен {1=1}
lRUCache.put(2, 2); //кэш равен {1=1, 2=2}
lRUCache.get(1);    //возвращаем 1
lRUCache.put(3, 3); //Ключ LRU был 2, вытесняет ключ 2, кеш равен {1=1, 3=3}
lRUCache.get(2);    //возвращает -1 (не найден)
lRUCache.put(4, 4); //Ключ LRU был 1, вытесняет ключ 1, кеш равен {4=4, 3=3}
lRUCache.get(1);    //возвращаем -1 (не найдено)
lRUCache.get(3);    //возвращаем 3
lRUCache.get(4);    //возвращаем 4

Ограничения:
1 <= capacity <= 3000
0 <= key <= 10^4
0 <= value <= 10^5
Для получения и размещения будет совершено максимум 2*10^5 вызовов.
*/

type node struct {
	key int
	val int
	prev *node
	next *node
}

type LRUCache struct {
    capacity int
	cache map[int]*node
	left *node
	right *node
}


// Constructor создаёт пустой LRU с двумя сторожевыми узлами.
// Время O(1), начальная память O(1): capacity не выделяется заранее.
// Хранимый кэш затем занимает O(C), где C=capacity; возвращаемый объект O(1).
// Все настоящие узлы лежат между left и right: слева LRU, справа MRU.
func Constructor(capacity int) LRUCache {
    cache := make(map[int]*node, 0)
	left, right := &node{}, &node{}
	// left.next=LRU, right.prev=самый недавний по использованию
	left.next, right.prev = right, left
	return LRUCache{capacity, cache, left, right}
}
// remove исключает настоящий узел из двусвязного списка за O(1) времени.
// Дополнительная память O(1); меняет связи соседей, но не map и не поля node.
// Предусловие: node уже в списке, его prev/next не nil; sentinels не удаляем.
func (this *LRUCache) remove(node *node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}
// insert добавляет отсоединённый узел перед right, на позицию MRU.
// Время O(1), дополнительная память O(1); меняет список и ссылки node.
// right.prev существует даже у пустого кэша: тогда это left.
func (this *LRUCache) insert(node *node) {
	prev, right := this.right.prev, this.right
	prev.next, right.prev = node, node
	node.prev = prev
	node.next = right
}

// Get возвращает значение и делает найденный ключ самым недавним.
// Среднее время O(1) для int-ключа map, дополнительная память O(1).
// Меняет порядок списка, но не число ключей и не их значения.
func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
		// Переставляем тот же узел в MRU; map продолжает ссылаться на него.
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1
}

// Put записывает значение, обновляет MRU и при переполнении удаляет LRU.
// Среднее амортизированное время O(1), рабочая память операции O(1),
// хранимые map/список O(C), где C=this.capacity; отдельного результата нет.
// Меняет значения, состав и порядок кэша. Даже обновление создаёт новый node.
func (this *LRUCache) Put(key int, value int) {
    // Старый узел ключа удаляем из списка до замены map-ссылки.
    // Иначе в списке осталось бы две записи одного ключа.
    if node, ok := this.cache[key]; ok {
		this.remove(node)
	}
	this.cache[key] = &node{key: key, val: value}
	this.insert(this.cache[key])
	// Временно возможны C+1 ключей. При положительном C удаляем настоящий
	// left.next; удаление и из списка, и из map восстанавливает их соответствие.
	if len(this.cache) > this.capacity {
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
