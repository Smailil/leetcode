package main

import "container/heap"

/*
Дан массив точек, где Points[i] = [x_i, y_i] представляет точку на плоскости XY и целое число k,
верните k ближайших точек к началу координат (0, 0).
Расстояние между двумя точками на плоскости XY — это евклидово расстояние (т. е. √(x1 — x2)^2 + (y1 — y2)^2).
Вы можете вернуть ответ в любом порядке.
Ответ гарантированно будет уникальным (за исключением порядка, в котором он находится).

Пример 1:

Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Объяснение:
Расстояние между (1, 3) и началом координат равно sqrt(10).
Расстояние между (-2, 2) и началом координат равно sqrt(8).
Поскольку sqrt(8) < sqrt(10), (-2, 2) ближе к началу координат.
Нам нужны только ближайшие k = 1 точки от начала координат, поэтому ответ — просто [[-2,2]].

Пример 2:
Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Пояснение: Ответ [[-2,4],[3,3]] также будет принят.

Ограничения:
1 <= k <= points.length <= 10^4
-10^4 <= x_i, y_i <= 10^4

*/

type MaxHeap [][3]int

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {return h[i][0] > h[j][0]}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}) {*h = append(*h, x.([3]int))}
func (h *MaxHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0:n-1]
    return x
}

func kClosest(points [][]int, k int) [][]int {
    mh := MaxHeap{}
    for _, point := range points {
		var p [3]int = [3]int{(point[0]*point[0]+point[1]*point[1]), point[0], point[1]}
        heap.Push(&mh, p)
        if len(mh) > k { heap.Pop(&mh) }
    }
    res := make([][]int, 0, len(mh))
	for _, m := range mh {
		res = append(res, []int{m[1], m[2]})
	}
	return res
}
