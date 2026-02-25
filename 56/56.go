package main

import "sort"

/*
Дан массив интервалов, где интервалы[i] = [starti, endi], объедините все перекрывающиеся интервалы и верните массив непересекающихся интервалов,
который охватывает все интервалы во входных данных.

Пример 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Пояснение: Поскольку интервалы [1,3] и [2,6] перекрываются, объедините их в [1,6].

Пример 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Пояснение: Интервалы [1,4] и [4,5] считаются перекрывающимися.

Пример 3:
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Пояснение: Интервалы [1,4] и [4,7] считаются перекрывающимися.

Ограничения:
1 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^4

*/

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
	
	output := make([][]int, 0, len(intervals))
    output = append(output, intervals[0])

    for _, interval := range intervals[1:] {
        start, end := interval[0], interval[1]
        lastEnd := output[len(output)-1][1]

        if start <= lastEnd {
            output[len(output)-1][1] = max(lastEnd, end)
        } else {
            output = append(output, []int{start, end})
        }
    }
    return output
}