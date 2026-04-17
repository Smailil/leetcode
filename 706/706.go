package main

/*
Спроектируйте HashMap без использования встроенных библиотек хеш-таблиц.

Реализуйте класс MyHashMap:

MyHashMap() инициализирует объект с пустой картой.
void put(int key, int value) вставляет пару (ключ, значение) в HashMap.
    Если ключ уже существует, обновляет соответствующее значение.
int get(int key) возвращает значение, с которым сопоставлен указанный ключ,
    или -1, если карта не содержит сопоставления для этого ключа.
void remove(key) удаляет ключ и его значение, если карта содержит это сопоставление.

Пример 1:

Вход
["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
[[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
Выход
[null, null, null, 1, -1, null, 1, null, -1]

Пояснение
MyHashMap myHashMap = new MyHashMap();
myHashMap.put(1, 1); // Карта теперь [[1,1]]
myHashMap.put(2, 2); // Карта теперь [[1,1], [2,2]]
myHashMap.get(1);    // возвращает 1, карта [[1,1], [2,2]]
myHashMap.get(3);    // возвращает -1 (не найдено), карта [[1,1], [2,2]]
myHashMap.put(2, 1); // Карта теперь [[1,1], [2,1]] (обновляет существующее значение)
myHashMap.get(2);    // возвращает 1, карта [[1,1], [2,1]]
myHashMap.remove(2); // удаляет сопоставление для 2, карта теперь [[1,1]]
myHashMap.get(2);    // возвращает -1 (не найдено), карта [[1,1]]

Ограничения:

0 <= key, value <= 10^6
Не более 10^4 вызовов put, get и remove.

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
