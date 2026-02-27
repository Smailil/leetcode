package main

import "sort"

/*
Дан массив временных интервалов встреч, где intervals[i] = [starti, endi], определите, может ли человек присутствовать на всех собраниях.

Пример 1:
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false

Пример 2:
Input: intervals = [[7,10],[2,4]]
Output: true

Constraints:
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti < endi <= 10^6

*/

type Interval struct {
   start int
   end   int
}

func canAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    for i := 1; i < len(intervals); i++ {
        if intervals[i-1].end > intervals[i].start {
            return false
        }
    }
    return true
}