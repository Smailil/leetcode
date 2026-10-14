package main

import "container/heap"

/*
Медиана — это значение в середине упорядоченного целочисленного списка.
Если размер списка четный, единственного среднего элемента нет,
а медиана — это среднее значение двух средних значений.
Например, для arr = [2,3,4] медиана равна 3.
Например, для arr = [2,3] медиана равна (2 + 3)/2 = 2,5.

Реализуйте класс MedianFinder:
MedianFinder() инициализирует объект MedianFinder.
void addNum(int num) добавляет целое число из потока данных в структуру данных.
double findMedian() возвращает медиану всех элементов на данный момент.
Принимаются ответы в пределах 10^-5 от фактического ответа.

Пример 1:
Вход:
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Выход:
[null, null, null, 1.5, null, 2.0]
Объяснение:
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // возвращает 1.5 (т.е. (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // возвращает 2.0

Ограничения:
-10^5 <= num <= 10^5
Перед вызовом findMedian в структуре данных будет хотя бы один элемент.
Будет выполнено не более 5 *10^4 вызовов addNum и findMedian.
*/

type Heap struct {
	Values   []int
	LessFunc func(int, int) bool
}

// Less делегирует сравнение LessFunc; здесь компараторы чистые и O(1).
// Рабочая память и булев результат O(1); h.Values не изменяется.
// Индексы i/j обеспечивает протокол container/heap.
func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
// Swap обменивает два элемента h.Values по индексам container/heap.
// Время и рабочая память O(1); порядок элементов h.Values изменяется.
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
// Len читает логическое число элементов за O(1) времени и памяти.
// Состояние не меняется, скалярный результат O(1).
func (h *Heap) Len() int           { return len(h.Values) }
// Peek читает корень непустой кучи за O(1) времени и памяти.
// Не меняет h; непустота обеспечивается инвариантом MedianFinder.
func (h *Heap) Peek() int          { return h.Values[0] }
// Pop — callback, снимающий последний элемент h.Values за O(1).
// container/heap предварительно переносит нужный корень в конец.
// Рабочая память и результат O(1), срез h.Values изменяется;
// вызов требует непустой кучи, общий буфер не освобождается срезом.
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
// Push — callback, добавляющий int в h.Values, а не полный heap.Push.
// Время O(1) амортизированно; перенос буфера может стоить O(t),
// t — число элементов. Рабочая память O(1) без буфера, рост до O(t+1).
// Изменяет h.Values; тип v.(int) обеспечен AddNum и собственным Pop.
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

// NewHeap создаёт пустую кучу с переданным направлением сравнения.
// Время, рабочая память и новый объект O(1); less пока не вызывается.
func NewHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

type MedianFinder struct {
	maxHeap *Heap  // для меньшей половины (максимальная куча)
	minHeap *Heap  // для большей половины (минимальная куча)
}

// Constructor создаёт max-кучу нижней и min-кучу верхней половины.
// Время, рабочая память и пустой результат O(1); внешних изменений нет.
// После вставок maxHeap имеет столько же элементов, сколько minHeap,
// или на один больше; каждый нижний элемент не больше любого верхнего.
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewHeap(func(a, b int) bool {
			return a > b  // корень - максимальный элемент
		}),
		minHeap: NewHeap(func(a, b int) bool {
			return a < b  // корень - минимальный элемент
		}),
	}
}

// AddNum сохраняет порядок половин, перенося их граничный элемент.
// Время O(log(t+1)) амортизированно, t — число прежних элементов;
// отдельный рост буфера может стоить O(t). Рабочая память O(1)
// без буферов, до O(t+1) при их переносе; хранимые данные O(t+1).
// Изменяет обе кучи mf; отдельного возвращаемого результата нет.
func (mf *MedianFinder) AddNum(num int) {
	// При равных размерах сначала добавляем в верхнюю половину,
	// затем её минимум переносим вниз: maxHeap станет на один больше.
	if (mf.maxHeap.Len()+mf.minHeap.Len())%2 == 0 {
		heap.Push(mf.minHeap, num)
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	} else {
		// Нижняя половина была на один больше: добавляем туда,
		// переносим её максимум вверх и снова уравниваем размеры.
		heap.Push(mf.maxHeap, num)
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	}
}

// FindMedian берёт один или два граничных корня упорядоченных половин.
// Время, рабочая память и скалярный результат O(1); mf не изменяется.
// По условию элементов хотя бы один: при нечётном размере maxHeap непуст,
// а при чётном непусты обе кучи.
func (mf *MedianFinder) FindMedian() float64 {
	if (mf.maxHeap.Len()+mf.minHeap.Len())%2 == 0 {
		// Преобразуем до сложения и деления: не теряется половина
		// нечётной суммы, а сложение не выполняется в типе int.
		return (float64(mf.maxHeap.Peek()) + float64(mf.minHeap.Peek())) / 2
	}
	return float64(mf.maxHeap.Peek())
}


/**
 * Объект MedianFinder будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
