package main

/*
Спроектируйте хэш-множество без использования встроенных библиотек хэш-таблиц.

Реализуйте класс MyHashSet:

void add(key) — вставляет значение key в хэш-множество.
bool contains(key) — возвращает true, если key присутствует в хэш-множестве, иначе false.
void remove(key) — удаляет key из хэш-множества. Если key не существует — ничего не делать.

Пример 1:

Входные данные
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Выходные данные
[null, null, null, true, false, null, true, null, false]

Пояснение
MyHashSet myHashSet = new MyHashSet();
myHashSet.add(1);      // множество = [1]
myHashSet.add(2);      // множество = [1, 2]
myHashSet.contains(1); // вернуть True
myHashSet.contains(3); // вернуть False (не найдено)
myHashSet.add(2);      // множество = [1, 2]
myHashSet.contains(2); // вернуть True
myHashSet.remove(2);   // множество = [1]
myHashSet.contains(2); // вернуть False (уже удалено)

Ограничения:

0 <= key <= 10^6
Не более 10^4 вызовов будет сделано к add, remove и contains.

*/

// ListNode представляет узел односвязного списка.
// Используется для организации цепочек в бакетах хеш-таблицы
// (разрешение коллизий методом цепочек / separate chaining).
type ListNode struct {
	key  int
	next *ListNode
}

// MyHashSet реализует множество с помощью хеш-таблицы.
// Коллизии обрабатываются с помощью отдельных связных списков.
type MyHashSet struct {
	set []*ListNode // Срез указателей на головные (фиктивные) узлы бакетов
}

// Constructor создает и инициализирует новое множество.
// Выделяет память под 10000 бакетов. Каждый бакет начинается с фиктивного узла,
// что значительно упрощает логику Add и Remove (не нужно отдельно обрабатывать
// случай, когда удаляемый/добавляемый элемент является первым в цепочке).
func Constructor() MyHashSet {
	set := make([]*ListNode, 10000)
	for i := range set {
		// Фиктивный (dummy) узел с нейтральным ключом.
		// Реальные данные всегда будут добавляться в cur.next
		set[i] = &ListNode{key: 0}
	}
	return MyHashSet{set: set}
}

// hash вычисляет индекс бакета для заданного ключа.
// Использует простую хеш-функцию на основе операции взятия остатка от деления.
func (this *MyHashSet) hash(key int) int {
	return key % len(this.set)
}

func (this *MyHashSet) Add(key int) {
	// Начинаем обход с фиктивного узла бакета
	cur := this.set[this.hash(key)]
	
	// Ищем конец цепочки или дубликат ключа
	for cur.next != nil {
		if cur.next.key == key {
			return // Ключ уже существует, ничего не делаем
		}
		cur = cur.next
	}
	// Добавляем новый узел в конец цепочки
	cur.next = &ListNode{key: key}
}


func (this *MyHashSet) Remove(key int) {
	cur := this.set[this.hash(key)]
	
	// Ищем узел, СЛЕДУЮЩИЙ ЗА КОТОРЫМ находится целевой ключ.
	// Благодаря фиктивному head-узлу мы всегда имеем доступ к предыдущему узлу.
	for cur.next != nil {
		if cur.next.key == key {
			// Исключаем узел из цепочки: предыдущий теперь указывает на следующий за удаляемым
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	cur := this.set[this.hash(key)]
	for cur.next != nil {
		if cur.next.key == key {
			return true
		}
		cur = cur.next
	}
	return false
}


/**
 * Объект MyHashSet создаётся и используется следующим образом:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
