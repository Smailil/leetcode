package main

/*
Разработайте и реализуйте структуру данных для кэша на основе алгоритма
LFU (Least Frequently Used — наименее часто используемый).

Реализуйте класс LFUCache:

LFUCache(int capacity) — инициализирует объект с заданной ёмкостью структуры данных.
int get(int key) — возвращает значение ключа, если он есть в кэше. Иначе возвращает -1.
void put(int key, int value) — обновляет значение ключа, если он присутствует,
или добавляет ключ, если его ещё нет. Когда кэш заполнен, перед вставкой нового элемента
необходимо вытеснить наименее часто используемый ключ.
При одинаковой частоте (т.е. два и более ключей с одинаковым счётчиком использования)
вытесняется наиболее давно использованный ключ.
Для определения LFU-ключа для каждого ключа в кэше ведётся счётчик обращений.
Ключ с наименьшим значением счётчика считается наименее часто используемым.

При первом добавлении ключа его счётчик устанавливается в 1 (благодаря операции put).
Счётчик увеличивается при каждом вызове get или put для этого ключа.

Операции get и put должны выполняться за O(1) амортизированного времени в среднем.



Пример 1:
Вход
["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
Выход
[null, null, null, 1, null, -1, 3, null, -1, 3, 4]

Пояснение
// cnt(x) = счётчик обращений для ключа x
// cache=[] показывает порядок последнего использования при равных счётчиках
//         (крайний левый элемент — использованный последним)
LFUCache lfu = new LFUCache(2);
lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
lfu.get(1);      // return 1
                 // cache=[1,2], cnt(2)=1, cnt(1)=2
lfu.put(3, 3);   // 2 — LFU-ключ, т.к. cnt(2)=1 минимально, вытесняем 2.
                 // cache=[3,1], cnt(3)=1, cnt(1)=2
lfu.get(2);      // return -1 (не найден)
lfu.get(3);      // return 3
                 // cache=[3,1], cnt(3)=2, cnt(1)=2
lfu.put(4, 4);   // У 1 и 3 одинаковый cnt, но 1 — LRU, вытесняем 1.
                 // cache=[4,3], cnt(4)=1, cnt(3)=2
lfu.get(1);      // return -1 (не найден)
lfu.get(3);      // return 3
                 // cache=[3,4], cnt(4)=1, cnt(3)=3
lfu.get(4);      // return 4
                 // cache=[4,3], cnt(4)=2, cnt(3)=3


Ограничения:
1 <= capacity <= 10^4
0 <= key <= 10^5
0 <= value <= 10^9
Не более 2 * 10^5 вызовов get и put.

*/

// LFUNode представляет узел двусвязного списка для хранения элемента кэша.
type LFUNode struct {
	key  int       // Ключ элемента
	val  int       // Значение элемента
	freq int       // Текущая частота обращений к элементу
	prev *LFUNode  // Указатель на предыдущий узел
	next *LFUNode  // Указатель на следующий узел
}

// LFUList представляет двусвязный список с фиктивными узлами (sentinel)
// для быстрого добавления/удаления элементов с одной и той же частотой.
type LFUList struct {
	left  *LFUNode // Фиктивный узел в начале списка
	right *LFUNode // Фиктивный узел в конце списка
	size  int      // Количество реальных узлов в списке
}

// NewLFUList создаёт и инициализирует пустой список LFUList.
func NewLFUList() *LFUList {
	left := &LFUNode{}
	right := &LFUNode{}
	left.next = right
	right.prev = left
	return &LFUList{
		left:  left,
		right: right,
		size:  0,
	}
}

// length возвращает текущее количество реальных элементов в списке.
func (ll *LFUList) length() int {
	return ll.size
}

// pushRight добавляет узел в конец списка (перед правым фиктивным узлом).
// Используется при добавлении нового элемента или перемещении в группу с большей частотой.
func (ll *LFUList) pushRight(node *LFUNode) {
	prev := ll.right.prev
	prev.next = node
	node.prev = prev
	node.next = ll.right
	ll.right.prev = node
	ll.size++
}

// pop удаляет указанный узел из списка, обновляя связи соседей.
func (ll *LFUList) pop(node *LFUNode) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
	// Обнуляем ссылки для предотвращения утечек памяти и побочных эффектов
	node.prev = nil
	node.next = nil
	ll.size--
}

// popLeft извлекает и возвращает первый реальный узел из начала списка.
// Используется для вытеснения наименее давно использованного элемента среди LFU.
func (ll *LFUList) popLeft() *LFUNode {
	node := ll.left.next
	ll.pop(node)
	return node
}

// LFUCache реализует кэш с политикой вытеснения наименее часто используемых элементов (LFU).
type LFUCache struct {
	capacity int                // Максимальная вместимость кэша
	lfuCount int                // Текущая минимальная частота обращений в кэше
	nodeMap  map[int]*LFUNode   // Быстрый доступ к узлу по ключу (O(1))
	listMap  map[int]*LFUList   // Группировка узлов по частоте обращений (O(1))
}

// Constructor создаёт новый экземпляр LFU-кэша заданной вместимости.
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		lfuCount: 0,
		nodeMap:  make(map[int]*LFUNode),
		listMap:  make(map[int]*LFUList),
	}
}

// counter увеличивает частоту обращения к узлу и перемещает его в список следующей частоты.
// Обновляет глобальный указатель lfuCount, если текущая группа частот стала пустой.
func (this *LFUCache) counter(node *LFUNode) {
	count := node.freq
	this.listMap[count].pop(node) // Удаляем узел из текущей частотной группы

	// Если удалили последний узел с минимальной частотой, сдвигаем указатель вправо
	if count == this.lfuCount && this.listMap[count].length() == 0 {
		this.lfuCount++
	}

	node.freq++ // Увеличиваем счётчик обращений
	// Если списка для новой частоты ещё нет, создаём его
	if _, exists := this.listMap[node.freq]; !exists {
		this.listMap[node.freq] = NewLFUList()
	}
	this.listMap[node.freq].pushRight(node) // Добавляем в конец новой группы (сохраняем порядок LRU внутри группы)
}

// Get возвращает значение по ключу. Если ключ отсутствует, возвращает -1.
// Автоматически обновляет частоту обращения к элементу.
func (this *LFUCache) Get(key int) int {
	node, exists := this.nodeMap[key]
	if !exists {
		return -1
	}
	this.counter(node)
	return node.val
}

// Put добавляет или обновляет пару ключ-значение в кэше.
// Если кэш заполнен, вытесняется элемент с минимальной частотой (при равенстве частот — наименее давно использованный).
func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	// Если ключ уже существует, обновляем значение и частоту
	if node, exists := this.nodeMap[key]; exists {
		node.val = value
		this.counter(node)
		return
	}

	// Если кэш заполнен, удаляем LFU/LRU элемент
	if len(this.nodeMap) == this.capacity {
		toRemove := this.listMap[this.lfuCount].popLeft()
		delete(this.nodeMap, toRemove.key)
	}

	// Создаём новый узел с начальной частотой 1
	node := &LFUNode{key: key, val: value, freq: 1}
	this.nodeMap[key] = node

	// Инициализируем список для частоты 1, если он ещё не создан
	if _, exists := this.listMap[1]; !exists {
		this.listMap[1] = NewLFUList()
	}
	this.listMap[1].pushRight(node)

	// При вставке нового элемента минимальная частота всегда сбрасывается на 1
	this.lfuCount = 1
}

/**
 * Объект LFUCache создаётся и используется следующим образом:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
