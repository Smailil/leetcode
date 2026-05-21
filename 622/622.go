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

type MyCircularQueue struct {

}


func Constructor(k int) MyCircularQueue {

}


func (this *MyCircularQueue) EnQueue(value int) bool {

}


func (this *MyCircularQueue) DeQueue() bool {

}


func (this *MyCircularQueue) Front() int {

}


func (this *MyCircularQueue) Rear() int {

}


func (this *MyCircularQueue) IsEmpty() bool {

}


func (this *MyCircularQueue) IsFull() bool {

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
