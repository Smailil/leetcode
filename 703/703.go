package main

/*
Вы работаете в приемной комиссии университета и вам необходимо отслеживать k-й наивысший результат 
тестов абитуриентов в режиме реального времени. Это помогает динамически определять пороговые баллы
для собеседований и поступления по мере того, как новые кандидаты представляют свои баллы.
Вам поручено реализовать класс, который для заданного целого числа k поддерживает поток результатов тестов
и постоянно возвращает k-й наивысший результат теста после подачи нового результата.
Точнее, мы ищем k-й наивысший балл в отсортированном списке всех баллов.

Реализуйте класс KthLargest:
KthLargest(int k, int[] nums) Инициализирует объект целым числом k и потоком чисел результатов тестов.
int add(int val) Добавляет в поток новое значение теста и возвращает элемент,
представляющий k-й по величине элемент в пуле результатов тестов на данный момент.

Пример 1:
Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output: [null, 4, 5, 5, 8, 8]
Объяснение:
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3); // return 4
kthLargest.add(5); // return 5
kthLargest.add(10); // return 5
kthLargest.add(9); // return 8
kthLargest.add(4); // return 8

Пример 2:
Input:
["KthLargest", "add", "add", "add", "add"]
[[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]
Output: [null, 7, 7, 7, 8]
Объяснение:
KthLargest kthLargest = new KthLargest(4, [7, 7, 7, 7, 8, 3]);
kthLargest.add(2); // return 7
kthLargest.add(10); // return 7
kthLargest.add(9); // return 7
kthLargest.add(9); // return 8

Ограничения:
0 <= nums.length <= 10^4
1 <= k <= nums.length + 1
-10^4 <= nums[i] <= 10^4
-10^4 <= val <= 10^4
Для добавления будет сделано не более 10^4 вызовов.

*/

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    n := h.Len()
    x := (*h)[n-1]
    *h = (*h)[0:n-1]
    return x
}

type KthLargest struct {
    k int
    minHeap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
    minHeap := &IntHeap{}
    heap.Init(minHeap)
    kthLargest := KthLargest{k, minHeap}
    for _, num := range nums {
        kthLargest.Add(num)
    }
    return kthLargest
}

func (this *KthLargest) Add(val int) int {
    if this.minHeap.Len() < this.k {
        heap.Push(this.minHeap, val)
    } else if val > (*this.minHeap)[0] {
        heap.Pop(this.minHeap)
        heap.Push(this.minHeap, val)
    }
    return (*this.minHeap)[0]
}