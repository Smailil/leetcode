package main

import "sort"

/*
Учитывая массив интервалов интервалов, где интервалы[i] = [starti, endi], верните минимальное количество интервалов, которое вам нужно удалить,
чтобы остальные интервалы не перекрывались.
Обратите внимание, что интервалы, соприкасающиеся только в одной точке, не перекрываются. Например, [1, 2] и [2, 3] не перекрываются.

Пример 1:
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Пояснение: [1,3] можно удалить, а остальные интервалы не перекрываться.

Пример 2:
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Пояснение: Вам нужно удалить два [1,2], чтобы остальные интервалы не перекрывались.

Пример 3:
Input: intervals = [[1,2],[2,3]]
Output: 0
Объяснение: Вам не нужно удалять какие-либо интервалы, поскольку они уже не перекрываются.

Ограничения:
1 <= intervals.length <= 10^5
intervals[i].length == 2
-5 * 10^4 <= starti < endi <= 5 * 10^4

*/

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    res := 0
    prevEnd := intervals[0][1]

    for _, interval := range intervals[1:] {
        start, end := interval[0], interval[1]
        if start >= prevEnd {
            prevEnd = end
        } else {
            res++
			if end < prevEnd {
				prevEnd = end
			}
        }
    }
    return res
}
