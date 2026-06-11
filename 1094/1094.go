package main

import (
	"container/heap"
	"sort"
)

/*
Есть автомобиль с `capacity` свободными местами.
Он едет только на восток, то есть не может развернуться и поехать на запад.

Даны целое число `capacity` и массив `trips`, где
`trips[i] = [numPassengersi, fromi, toi]` означает,
что в `i`-й поездке нужно забрать `numPassengersi` пассажиров
в точке `fromi` и высадить их в точке `toi`.
Позиции заданы как количество километров на восток
от начального положения автомобиля.

Верните `true`, если можно забрать и высадить всех пассажиров
для всех заданных поездок, иначе верните `false`.

Пример 1:
Вход: trips = [[2,1,5],[3,3,7]], capacity = 4
Выход: false

Пример 2:
Вход: trips = [[2,1,5],[3,3,7]], capacity = 5
Выход: true

Ограничения:
1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi <= 100
0 <= fromi < toi <= 1000
1 <= capacity <= 10^5
*/

func carPooling(trips [][]int, capacity int) bool {
	// Сортируем поездки по точке посадки (start), чтобы обрабатывать их в хронологическом порядке.
	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	// MinHeap хранит активные поездки в формате [end, numPassengers].
	// Используется для быстрого поиска поездки с самой ранней точкой высадки.
	h := &MinHeap{}
	heap.Init(h)
	curPass := 0 // Текущее количество пассажиров в машине.

	for _, trip := range trips {
		numPass, start, end := trip[0], trip[1], trip[2]

		// Высаживаем всех пассажиров, чьи поездки завершились до или в текущей точке посадки.
		// (*h)[0][0] — это минимальная точка высадки на вершине кучи.
		for h.Len() > 0 && (*h)[0][0] <= start {
			curPass -= heap.Pop(h).([2]int)[1]
		}

		// Сажаем новых пассажиров.
		curPass += numPass

		// Если текущее количество пассажиров превышает вместимость — вернуть false.
		if curPass > capacity {
			return false
		}

		// Добавляем текущую поездку в кучу: запоминаем точку высадки и количество пассажиров.
		heap.Push(h, [2]int{end, numPass})
	}

	// Все поездки обработаны, вместимость ни разу не превышена.
	return true
}

// MinHeap — минимальная куча, упорядоченная по точке высадки [0].
// Позволяет за O(log n) извлекать поездку с самым ранним временем высадки.
type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] } // Сравниваем по точке высадки.
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]          // Забираем последний элемент (heap гарантирует, что он минимальный после heap.Pop).
	*h = old[0 : n-1]      // Уменьшаем размер кучи.
	return x
}
