package main

/*
Вам дана сеть из n узлов, помеченных от 1 до n. Вам также дано время, список времени в пути по направлению ребра times[i] = (ui, vi, wi), 
где ui — исходный узел, vi — целевой узел, а wi — время, необходимое сигналу для прохождения от источника к цели.

Мы отправим сигнал из данного узла k. Возвращает минимальное время, необходимое всем n узлам для получения сигнала. 
Если все n узлов не могут получить сигнал, верните -1.

Пример 1:
https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png
Input: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2

Пример 2:
Input: times = [[1,2,1]], n = 2, k = 1
Output: 1

Пример 3:
Input: times = [[1,2,1]], n = 2, k = 2
Output: -1

Ограничения:
1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= ui, vi <= n
ui != vi
0 <= wi <= 100
Все пары (ui, vi) уникальны. (т. е. отсутствие кратных ребер.)

*/

import (
    "container/heap"
)

// Структура для элементов кучи
type Item struct {
    node     int
    distance int
}

// Реализация интерфейса heap.Interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    *pq = old[:n-1]
    return item
}

func networkDelayTime(times [][]int, n int, k int) int {
    // Создание графа (adjacency list)
    edges := make(map[int][][2]int)
    for _, time := range times {
        u, v, w := time[0], time[1], time[2]
        edges[u] = append(edges[u], [2]int{v, w})
    }

    // Инициализация приоритетной очереди (min-heap)
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    heap.Push(&pq, &Item{node: k, distance: 0})

    // Множество посещённых узлов
    visit := make(map[int]bool)
    t := 0

    // Алгоритм Дейкстры
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        w1, n1 := item.distance, item.node

        if visit[n1] {
            continue
        }
        visit[n1] = true
        t = w1

        for _, edge := range edges[n1] {
            n2, w2 := edge[0], edge[1]
            if !visit[n2] {
                heap.Push(&pq, &Item{node: n2, distance: w1 + w2})
            }
        }
    }

    if len(visit) == n {
        return t
    }
    return -1
}
