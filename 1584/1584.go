package main

/*
Вам дан массив точек, представляющий целочисленные координаты некоторых точек на 2D-плоскости, где points[i] = [xi, yi].
Стоимость соединения двух точек [xi, yi] и [xj, yj] равна манхэттенскому расстоянию между ними: |xi -xj| + |yi -yj|, где |val| обозначает абсолютное значение val.
Верните минимальную стоимость соединения всех точек. Все точки связаны, если между любыми двумя точками существует ровно один простой путь.

Пример 1:
https://assets.leetcode.com/uploads/2020/08/26/d.png
Input: points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
Output: 20
Объяснение: 
https://assets.leetcode.com/uploads/2020/08/26/c.png
Мы можем соединить точки, как показано выше, чтобы получить минимальную стоимость 20.
Обратите внимание, что между каждой парой точек существует уникальный путь.

Пример 2:
Input: points = [[3,12],[-2,5],[-4,1]]
Output: 18
 
Ограничения:
1 <= points.length <= 1000
-10^6 <= xi, yi <= 10^6
Все пары (xi, yi) различны.

*/

import (
	// "container/heap"
	"math"
)

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	
	// minDist[i] хранит минимальное расстояние от узла i до посещённого множества
	minDist := make([]int, n)
	for i := range minDist {
		minDist[i] = math.MaxInt32
	}
	minDist[0] = 0
	
	// visit[i] = true, если узел i уже посещён
	visit := make([]bool, n)
	
	res := 0
	nodesVisited := 0
	
	for nodesVisited < n {
		// Находим непосещённый узел с минимальным расстоянием (O(n))
		minCost := math.MaxInt32
		currNode := -1
		
		for i := range n {
			if !visit[i] && minDist[i] < minCost {
				minCost = minDist[i]
				currNode = i
			}
		}
		
		// Добавляем узел к посещённым
		visit[currNode] = true
		res += minCost
		nodesVisited++
		
		// Обновляем расстояния до всех непосещённых соседей (O(n))
		for nextNode := range n {
			if !visit[nextNode] {
				dist := int(math.Abs(float64(points[currNode][0]-points[nextNode][0])) +
				            math.Abs(float64(points[currNode][1]-points[nextNode][1])))
				if dist < minDist[nextNode] {
					minDist[nextNode] = dist
				}
			}
		}
	}
	
	return res
}

// // Edge представляет ребро графа: стоимость и индекс узла
// type Edge struct {
// 	cost int
// 	node int
// }

// // MinHeap реализует интерфейс heap.Interface для приоритетной очереди
// type MinHeap []Edge

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(Edge))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func minCostConnectPoints(points [][]int) int {
// 	n := len(points)
	
// 	// Создаём список смежности
// 	adj := make([][]Edge, n)
// 	for i := 0; i < n; i++ {
// 		x1, y1 := points[i][0], points[i][1]
// 		for j := i + 1; j < n; j++ {
// 			x2, y2 := points[j][0], points[j][1]
// 			dist := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
// 			adj[i] = append(adj[i], Edge{cost: dist, node: j})
// 			adj[j] = append(adj[j], Edge{cost: dist, node: i})
// 		}
// 	}

// 	// Алгоритм Прима
// 	res := 0
// 	visit := make(map[int]bool)
// 	minH := &MinHeap{}
// 	heap.Init(minH)
// 	heap.Push(minH, Edge{cost: 0, node: 0})

// 	for len(visit) < n {
// 		edge := heap.Pop(minH).(Edge)
// 		cost, i := edge.cost, edge.node
		
// 		if visit[i] {
// 			continue
// 		}
		
// 		res += cost
// 		visit[i] = true
		
// 		for _, nei := range adj[i] {
// 			if !visit[nei.node] {
// 				heap.Push(minH, nei)
// 			}
// 		}
// 	}

// 	return res
// }
