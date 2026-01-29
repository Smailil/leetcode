package main

import "container/heap"

/*
Медиана — это среднее значение в упорядоченном целочисленном списке.
Если размер списка четный, среднего значения нет, а медиана — это среднее значение двух средних значений.
Например, для arr = [2,3,4] медиана равна 3.
Например, для arr = [2,3] медиана равна (2 + 3)/2 = 2,5.

Реализуйте класс MedianFinder:

MedianFinder() инициализирует объект MedianFinder.

void addNum(int num) добавляет целое число из потока данных в структуру данных.

double findMedian() возвращает медиану всех элементов на данный момент.
Принимаются ответы в пределах 10^-5 от фактического ответа.

Пример 1:
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]
Объяснение
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (т.е. (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0

Ограничения:
-10^5 <= num <= 10^5
Перед вызовом findMedian в структуре данных будет хотя бы один элемент.
Будет выполнено не более 5 *10^4 вызовов addNum и findMedian.

*/

type Heap struct {
	Values   []int
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() int          { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

func NewHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

type MedianFinder struct {
	maxHeap *Heap  // для меньшей половины (максимальная куча)
	minHeap *Heap  // для большей половины (минимальная куча)
}

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

func (mf *MedianFinder) AddNum(num int) {
	if (mf.maxHeap.Len()+mf.minHeap.Len())%2 == 0 {
		heap.Push(mf.minHeap, num)
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	} else {
		heap.Push(mf.maxHeap, num)
		heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if (mf.maxHeap.Len()+mf.minHeap.Len())%2 == 0 {
		return (float64(mf.maxHeap.Peek()) + float64(mf.minHeap.Peek())) / 2
	}
	return float64(mf.maxHeap.Peek())
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
