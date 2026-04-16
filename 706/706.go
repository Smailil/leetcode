package main

/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
 

Example 1:

Input
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Output
[null, null, null, 1, -1, null, 1, null, -1]

Explanation
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // The map is now [[1,1]]
myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
 

Constraints:

0 <= key, value <= 106
At most 104 calls will be made to put, get, and remove.

*/

// HashNode представляет узел связного списка, используемого для разрешения коллизий в хеш-таблице (метод цепочек).
type HashNode struct {
	key, val int
	next     *HashNode
}

type MyHashMap struct {
	data []*HashNode // Срез указателей на dummy-головы каждого бакета
}

func Constructor() MyHashMap {
	data := make([]*HashNode, 1000)
	for i := range data {
		data[i] = &HashNode{key: -1, val: -1} // Dummy-узел для удобства операций
	}
	return MyHashMap{data: data}
}

// hash вычисляет индекс бакета для заданного ключа.
// Используется простая хеш-функция на основе операции взятия остатка от деления на размер таблицы.
func (this *MyHashMap) hash(key int) int {
	return key % len(this.data)
}

// Put вставляет или обновляет пару (ключ, значение) в хеш-таблице.
func (this *MyHashMap) Put(key int, value int) {
	// Получаем dummy-голову соответствующего бакета
	cur := this.data[this.hash(key)]
	
	// Проходим по списку, проверяя следующий узел
	for cur.next != nil {
		if cur.next.key == key {
			// Ключ уже существует: обновляем значение и выходим
			cur.next.val = value
			return
		}
		cur = cur.next
	}
	
	// Ключ не найден в цепочке: добавляем новый узел в конец списка
	cur.next = &HashNode{key: key, val: value}
}

// Get возвращает значение, связанное с указанным ключом.
// Если ключ отсутствует в таблице, возвращает -1.
func (this *MyHashMap) Get(key int) int {
	// Начинаем поиск сразу с первого реального элемента бакета (после dummy-головы)
	cur := this.data[this.hash(key)].next
	for cur != nil {
		if cur.key == key {
			return cur.val // Ключ найден, возвращаем значение
		}
		cur = cur.next
	}
	return -1 // Ключ не найден
}

// Remove удаляет пару с указанным ключом из хеш-таблицы.
// Если ключ не найден, метод завершается без изменений.
func (this *MyHashMap) Remove(key int) {
	cur := this.data[this.hash(key)]
	
	// Ищем узел, следующий за которым находится целевой ключ
	for cur.next != nil {
		if cur.next.key == key {
			// Ключ найден: исключаем следующий узел из цепочки, перенаправляя указатель
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
