package main

/*
Спроектируйте структуру данных, которая соответствует ограничениям кэша наименее недавно используемого (LRU).

Реализуйте класс LRUCache:
LRUCache(int емкость) Инициализировать кэш LRU с емкостью положительного размера.
int get(int key) Возвращает значение ключа, если ключ существует, в противном случае возвращает -1.
void put(int key, int value) Обновить значение ключа, если ключ существует.
В противном случае добавьте пару ключ-значение в кеш. Если количество ключей превышает емкость этой операции, удалите ключ, который использовался реже всего.
Каждая из функций get и put должна выполняться со средней временной сложностью O(1).

Пример 1:
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]
Объяснение
LRUCache lRUCache = новый LRUCache(2);
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


func Constructor(capacity int) LRUCache {
    cache := make(map[int]*node, 0)
	left, right := &node{}, &node{}
	// left.next=LRU, right.prev=самый недавний по использованию
	left.next, right.prev = right, left
	return LRUCache{capacity, cache, left, right}
}
// удалить из списка
func (this *LRUCache) remove(node *node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}
// вставить узел справа
func (this *LRUCache) insert(node *node) {
	prev, right := this.right.prev, this.right
	prev.next, right.prev = node, node
	node.prev = prev
	node.next = right
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
		this.remove(node)
		this.insert(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
		this.remove(node)
	}
	this.cache[key] = &node{key: key, val: value}
	this.insert(this.cache[key])
	if len(this.cache) > this.capacity {
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
