package main

import "sort"

/*
Дан массив временных интервалов встреч,
где intervals[i] = [start_i, end_i].
Определите, может ли человек присутствовать на всех собраниях.

Пример 1:
Вход: intervals = [[0,30],[5,10],[15,20]]
Выход: false

Пример 2:
Вход: intervals = [[7,10],[2,4]]
Выход: true

Ограничения:
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= start_i < end_i <= 10^6
*/

type Interval struct {
   start int
   end   int
}

// canAttendMeetings сортирует встречи по началу и проверяет соседние пары.
// Время O(n*log(n+1)), n=len(intervals), с последующим линейным проходом.
// Рабочая память O(log(n+1)) для сортировки; результат bool O(1).
// Порядок элементов входного intervals меняется на месте.
func canAttendMeetings(intervals []Interval) bool {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i].start < intervals[j].start
    })

    // Если соседние пары не пересекаются, более далёкие тоже не пересекутся.
    // Начало с 1 защищает i-1; пустой вход даёт true.
    for i := 1; i < len(intervals); i++ {
        // Окончание равно началу — допустимый переход без наложения.
        if intervals[i-1].end > intervals[i].start {
            return false
        }
    }
    return true
}