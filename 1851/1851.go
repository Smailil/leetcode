package main

/*
Вам дан двумерный целочисленный массив интервалов, где интервалы[i] = [lefti, righti] описывают i-й интервал, начинающийся с lefti и заканчивающийся righti (включительно). 
Размер интервала определяется как количество содержащихся в нем целых чисел или, более формально, righti -lefti + 1.
Вам также предоставляются запросы целочисленного массива. Ответом на j-й запрос является размер наименьшего интервала i такого, что lefti <= query[j] <= righti. 
Если такого интервала не существует, ответ – -1.
Возвращает массив, содержащий ответы на запросы.

Пример 1:
Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
Output: [3,3,1,4]
Объяснение: Запросы обрабатываются следующим образом:
-Запрос = 2: интервал [2,4] — это наименьший интервал, содержащий 2. Ответ: 4 — 2 + 1 = 3.
-Запрос = 3: Интервал [2,4] — это наименьший интервал, содержащий 3. Ответ: 4 — 2 + 1 = 3.
-Запрос = 4: Интервал [4,4] — это наименьший интервал, содержащий 4. Ответ: 4 — 4 + 1 = 1.
-Запрос = 5: Интервал [3,6] — это наименьший интервал, содержащий 5. Ответ: 6 — 3 + 1 = 4.

Пример 2:
Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
Output: [2,-1,4,6]
Объяснение: Запросы обрабатываются следующим образом:
-Запрос = 2: Интервал [2,3] — это наименьший интервал, содержащий 2. Ответ: 3 — 2 + 1 = 2.
-Запрос = 19: Ни один из интервалов не содержит 19. Ответ --1.
-Запрос = 5: Интервал [2,5] — это наименьший интервал, содержащий 5. Ответ: 5 — 2 + 1 = 4.
-Запрос = 22: Интервал [20,25] — это наименьший интервал, содержащий 22. Ответ: 25 — 20 + 1 = 6.

Ограничения:
1 <= intervals.length <= 10^5
1 <= queries.length <= 10^5
intervals[i].length == 2
1 <= lefti <= righti <= 10^7
1 <= queries[j] <= 10^7

*/

import (
	"container/heap"
	"sort"
)

// Interval представляет элемент кучи: размер интервала и его правый конец
type Interval struct {
	size int // r - l + 1
	end  int // r
}

// MinHeap реализует интерфейс heap.Interface для минимальной кучи
type MinHeap []Interval

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].size < h[j].size } // минимальная куча по размеру
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	// Сортируем интервалы по началу
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Создаём копию запросов для сортировки
	sortedQueries := make([]int, len(queries))
	copy(sortedQueries, queries)
	sort.Ints(sortedQueries)

	// Результат для каждого запроса
	res := make(map[int]int, len(queries))
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	i := 0
	for _, q := range sortedQueries {
		// Добавляем все интервалы, которые начинаются до или в момент запроса
		for i < len(intervals) && intervals[i][0] <= q {
			l, r := intervals[i][0], intervals[i][1]
			heap.Push(minHeap, Interval{size: r - l + 1, end: r})
			i++
		}

		// Удаляем интервалы, которые заканчиваются до запроса
		for minHeap.Len() > 0 && (*minHeap)[0].end < q {
			heap.Pop(minHeap)
		}

		// Записываем результат
		if minHeap.Len() > 0 {
			res[q] = (*minHeap)[0].size
		} else {
			res[q] = -1
		}
	}

	// Возвращаем результаты в оригинальном порядке запросов
	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = res[q]
	}

	return result
}
