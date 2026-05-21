package main

/*
Спроектируйте реализацию кольцевой очереди.
Кольцевая очередь — линейная структура данных, в которой операции выполняются
по принципу FIFO (первым пришёл — первым вышел), а последняя позиция
соединена с первой, образуя круг. Её также называют «кольцевым буфером».

Одно из преимуществ кольцевой очереди — возможность использовать
освободившееся место в начале. В обычной очереди после заполнения
вставить новый элемент невозможно, даже если место спереди освободилось.
Кольцевая очередь позволяет использовать это место для новых значений.

Реализуйте класс MyCircularQueue:

MyCircularQueue(k) — инициализирует объект с размером очереди k.
int Front() — возвращает первый элемент очереди. Если очередь пуста — вернуть -1.
int Rear() — возвращает последний элемент очереди. Если очередь пуста — вернуть -1.
boolean enQueue(int value) — вставляет элемент в кольцевую очередь.
  Вернуть true, если операция выполнена успешно.
boolean deQueue() — удаляет элемент из кольцевой очереди.
  Вернуть true, если операция выполнена успешно.
boolean isEmpty() — проверяет, пуста ли кольцевая очередь.
boolean isFull() — проверяет, заполнена ли кольцевая очередь.
Задачу необходимо решить без использования встроенной структуры данных очереди.


Пример 1:
Ввод
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
Вывод
[null, true, true, true, false, 3, true, true, true, 4]

Пояснение
MyCircularQueue myCircularQueue = new MyCircularQueue(3);
myCircularQueue.enQueue(1); // вернуть True
myCircularQueue.enQueue(2); // вернуть True
myCircularQueue.enQueue(3); // вернуть True
myCircularQueue.enQueue(4); // вернуть False
myCircularQueue.Rear();     // вернуть 3
myCircularQueue.isFull();   // вернуть True
myCircularQueue.deQueue();  // вернуть True
myCircularQueue.enQueue(4); // вернуть True
myCircularQueue.Rear();     // вернуть 4


Ограничения:
1 <= k <= 1000
0 <= value <= 1000
Суммарно не более 3000 вызовов enQueue, deQueue, Front, Rear, isEmpty и isFull.

*/
// MyListNode представляет узел двусвязного списка.
type MyListNode struct {
    val  int       // Значение, хранящееся в узле
    next *MyListNode // Указатель на следующий узел
    prev *MyListNode // Указатель на предыдущий узел
}

// MyCircularQueue реализует кольцевую очередь на базе двусвязного списка.
// Использование фиктивных (sentinel) узлов left и right упрощает граничные условия вставки/удаления.
type MyCircularQueue struct {
    space int       // Количество свободных мест в очереди
    left  *MyListNode // Фиктивный узел-голова (перед ним элементов нет)
    right *MyListNode // Фиктивный узел-хвост (после него элементов нет)
}

// Constructor инициализирует очередь с максимальной вместимостью k.
func Constructor(k int) MyCircularQueue {
    // Создаём два фиктивных узла и связываем их между собой
    left := &MyListNode{val: 0}
    right := &MyListNode{val: 0, prev: left}
    left.next = right

    return MyCircularQueue{
        space: k,
        left:  left,
        right: right,
    }
}

// EnQueue добавляет элемент в конец очереди.
// Возвращает true при успешном добавлении, false если очередь уже заполнена.
func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }

    // Создаём новый узел и вставляем его перед right-ограничителем
    cur := &MyListNode{val: value, next: this.right, prev: this.right.prev}
    // Обновляем указатели соседей для включения нового узла в список
    this.right.prev.next = cur
    this.right.prev = cur

    this.space-- // Уменьшаем счётчик доступных мест
    return true
}

// DeQueue удаляет элемент из начала очереди.
// Возвращает true при успешном удалении, false если очередь пуста.
func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }

    // "Пропускаем" первый реальный узел, соединяя left сразу со вторым элементом
    this.left.next = this.left.next.next
    // Обратный указатель второго элемента теперь должен указывать на left
    this.left.next.prev = this.left

    this.space++ // Освобождаем одно место в очереди
    return true
}

// Front возвращает значение первого элемента очереди.
// Если очередь пуста, возвращает -1.
func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    // Первый реальный элемент всегда находится сразу после left
    return this.left.next.val
}

// Rear возвращает значение последнего элемента очереди.
// Если очередь пуста, возвращает -1.
func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    // Последний реальный элемент всегда находится сразу перед right
    return this.right.prev.val
}

// IsEmpty проверяет, пуста ли очередь.
// Очередь пуста, когда left и right указывают друг на друга.
func (this *MyCircularQueue) IsEmpty() bool {
    return this.left.next == this.right
}

// IsFull проверяет, достигла ли очередь своей максимальной вместимости.
func (this *MyCircularQueue) IsFull() bool {
    return this.space == 0
}


/**
 * Объект MyCircularQueue создаётся и используется следующим образом:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
