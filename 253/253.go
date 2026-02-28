package main

import "sort"

/*
Дан массив интервалов времени встреч, где интервалы [i] = [starti, endi], верните минимальное количество требуемых конференц-залов.

Пример 1:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2

Пример 2:
Input: intervals = [[7,10],[2,4]]
Output: 1

Ограничения:
1 <= intervals.length <= 10^4
0 <= starti < endi <= 10^6

*/

type Interval struct {
   start int
   end   int
}

func minMeetingRooms(intervals []Interval) int {
    start := make([]int, len(intervals))
    end := make([]int, len(intervals))

    for i, interval := range intervals {
        start[i] = interval.start
        end[i] = interval.end
    }

    sort.Ints(start)
    sort.Ints(end)

    maxRooms, countRooms := 0, 0
    startIdx, endIdx := 0, 0

    for startIdx < len(intervals) {
        if start[startIdx] < end[endIdx] {
            startIdx++
            countRooms++
        } else {
            endIdx++
            countRooms--
        }
        if countRooms > maxRooms {
            maxRooms = countRooms
        }
    }

    return maxRooms
}
