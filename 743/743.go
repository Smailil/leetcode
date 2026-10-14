package main

/*
Дана сеть из n узлов, пронумерованных от 1 до n.
Также дан список времён прохождения по направленным рёбрам:
times[i] = (u_i, v_i, w_i), где u_i — исходный узел, v_i — целевой узел,
а w_i — время прохождения сигнала от исходного узла к целевому.

Сигнал отправляется из заданного узла k.
Верните минимальное время, за которое все n узлов получат сигнал.
Если не все n узлов могут получить сигнал, верните -1.

Пример 1:
https://assets.leetcode.com/uploads/2019/05/23/931_example_1.png
Вход: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Выход: 2

Пример 2:
Вход: times = [[1,2,1]], n = 2, k = 1
Выход: 1

Пример 3:
Вход: times = [[1,2,1]], n = 2, k = 2
Выход: -1

Ограничения:
1 <= k <= n <= 100
1 <= times.length <= 6000
times[i].length == 3
1 <= u_i, v_i <= n
u_i != v_i
0 <= w_i <= 100
Все пары (u_i, v_i) уникальны, то есть кратных рёбер нет.
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

// Len даёт размер за O(1) времени и памяти; ничего не меняет.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less сравнивает расстояния для min-heap; время/память/результат O(1).
// Индексы корректны по контракту heap, элементы — ненулевые *Item.
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].distance < pq[j].distance
}

// Swap меняет элементы общего среза; время и рабочая память O(1).
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

// Push добавляет *Item: время O(1) амортизированно, O(q) при росте буфера.
// Рабочая память при росте O(q), q — размер кучи; меняет очередь.
// Это callback, не вся операция heap за O(log(q+1)) амортизированно.
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    *pq = append(*pq, item)
}

// Pop снимает уже переставленный heap элемент с конца непустой очереди.
// Время, рабочая память и возвращаемая ссылка O(1); очередь укорачивается.
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    // Убираем удерживаемую ссылку в свободном слоте backing-массива.
    old[n-1] = nil
    *pq = old[:n-1]
    return item
}

// networkDelayTime применяет Дейкстру: узел фиксируется при первом снятии
// минимального расстояния; неотрицательные веса делают его окончательным.
// Среднее время O(n+E*log(E+1)), E=len(times); рабочая память O(n+E).
// Heap хранит дубли узлов, до E+1 записей; результат int O(1).
// Вход times не меняется; индексы и тройки рёбер верны по условию.
func networkDelayTime(times [][]int, n int, k int) int {
    // Создание графа (списка смежности)
    edges := make(map[int][][2]int)
    for _, time := range times {
        u, v, w := time[0], time[1], time[2]
        edges[u] = append(edges[u], [2]int{v, w})
    }

    // Инициализация приоритетной очереди (минимальной кучи)
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

        // Прежняя извлечённая запись этого узла уже закрепила лучший путь.
        if visit[n1] {
            continue
        }
        visit[n1] = true
        // Окончательные расстояния не убывают; последнее — их максимум.
        t = w1

        for _, edge := range edges[n1] {
            n2, w2 := edge[0], edge[1]
            // Ребро обрабатывается из окончательно выбранного узла один раз.
            // Ещё не выбранный сосед может иметь несколько кандидатов в куче.
            if !visit[n2] {
                heap.Push(&pq, &Item{node: n2, distance: w1 + w2})
            }
        }
    }

    // Максимум кратчайших расстояний — время сигнала для всех узлов.
    // Недостижимого узла не будет в visit.
    if len(visit) == n {
        return t
    }
    return -1
}
