package main

import "container/heap"

/*
Вам дан массив задач ЦП, каждая из которых помечена буквой от A до Z и числом n.
Каждый интервал ЦП может быть бездействующим или разрешать выполнение одной задачи.
Задачи можно выполнять в любом порядке, но есть ограничение: между двумя задачами
с одинаковой меткой должен быть промежуток не менее n интервалов.
Возвращает минимальное количество интервалов ЦП, необходимое для выполнения всех задач.

Пример 1:
Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Пояснение: Возможная последовательность: A -> B -> ожидание -> A -> B -> ожидание -> A -> B.
После выполнения задачи A вы должны подождать два интервала, прежде чем снова выполнить A.
То же самое относится и к задаче B.
В 3-м интервале ни A, ни B сделать невозможно, поэтому вы бездействуете.
К 4-му интервалу можно снова выполнить A, так как прошло 2 интервала.

Пример 2:
Input: tasks = ["A","C","A","B","D","B"], n = 1
Output: 6
Пояснение: Возможная последовательность: A -> B -> C -> D -> A -> B.
Если интервал охлаждения равен 1, вы можете повторить задачу только после одной другой задачи.

Пример 3:
Input: tasks = ["A","A","A", "B","B","B"], n = 3
Output: 10
Пояснение: Возможная последовательность: A -> B -> простой -> простой -> A -> B -> простой -> простой -> A -> B.
Есть всего два типа заданий A и B, которые нужно разделить 3 интервалами.
Это приводит к двойному простою между повторениями этих задач.

Ограничения:
1 <= tasks.length <= 10^4
tasks[i] — это заглавная английская буква.
0 <= n <= 100

*/

type MaxHeap []int	

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MaxHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0:n-1]
    return x
}

type Queue [][2]int

func (q *Queue) Push(x [2]int) {*q = append(*q, x)}
func (q *Queue) Pop() [2]int {
    x := (*q)[0]
    *q = (*q)[1:]
    return x
}

func leastInterval(tasks []byte, n int) int {
    count := make(map[byte]int, 26)
	for _, task := range tasks {
		count[task]++
	}
	mh := make(MaxHeap, 0, 26)
	for _, val := range count {
		heap.Push(&mh, val)
	}
	q := make(Queue, 0, 26)
	time := 0
	for len(mh) > 0 || len(q) > 0 {
		time++
		if len(mh) > 0 {
			val := heap.Pop(&mh).(int) - 1
			if val != 0 {
				q.Push([2]int{val, time+n})
			}
		}
		if len(q) > 0 {
			if q[0][1] == time {
				heap.Push(&mh, q.Pop()[0])
			}
		}
	}
	return time
}
