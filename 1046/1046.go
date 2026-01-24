package main

import "container/heap"

/*
Вам дан массив целых чисел, где Stones[i] — это вес i-го камня.
Мы играем в игру с камнями. На каждом ходу мы выбираем два самых тяжелых камня и разбиваем их вместе.
Предположим, что два самых тяжелых камня имеют вес x и y, причем x <= y. Результат этого удара:
Если x == y, оба камня уничтожаются, и
Если x != y, камень веса x уничтожается, а камень веса y приобретает новый вес y -x.
В конце игры остается не более одного камня.
Верните вес последнего оставшегося камня. Если камней не осталось, верните 0.

Пример 1:
Input: stones = [2,7,4,1,8,1]
Output: 1
Объяснение:
Мы объединяем 7 и 8, чтобы получить 1, поэтому массив преобразуется в [2,4,1,1,1], а затем:
мы объединяем 2 и 4, чтобы получить 2, поэтому массив преобразуется в [2,1,1,1], затем
мы объединяем 2 и 1, чтобы получить 1, поэтому массив преобразуется в [1,1,1], затем
мы объединяем 1 и 1, чтобы получить 0, поэтому массив преобразуется в [1], тогда это значение последнего камня.

Пример 2:
Input: stones = [1]
Output: 1

Ограничения:
1 <= stones.length <= 30
1 <= stones[i] <= 1000

*/

type IntHeap []int

func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *IntHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0:n-1]
    return x
}

func lastStoneWeight(stones []int) int {
    ih := IntHeap(stones)
    heap.Init(&ih)
    for ih.Len() > 1 {
        x, y := heap.Pop(&ih).(int), heap.Pop(&ih).(int)
        if x != y {
            heap.Push(&ih, x-y)
        }
    }
    
    if ih.Len() == 0 {
        return 0
    }
    
    return heap.Pop(&ih).(int)
}
