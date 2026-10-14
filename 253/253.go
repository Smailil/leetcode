package main

import "sort"

/*
Дан массив интервалов времени встреч,
где intervals[i] = [start_i, end_i].
Верните минимальное количество требуемых конференц-залов.

Пример 1:
Вход: intervals = [[0,30],[5,10],[15,20]]
Выход: 2

Пример 2:
Вход: intervals = [[7,10],[2,4]]
Выход: 1

Ограничения:
1 <= intervals.length <= 10^4
0 <= start_i < end_i <= 10^6
*/

type Interval struct {
   start int
   end   int
}

// minMeetingRooms считает пиковое число активных встреч проходом по событиям.
// Начала и концы сортируются отдельно, чтобы обрабатывать их по времени.
// Время O(n*log(n+1)), n=len(intervals); рабочая память O(n).
// Результат int занимает O(1); вход не меняется, сортируются новые массивы.
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

    // countRooms — число начатых, но ещё не завершённых встреч.
    // Пока есть необработанное начало, есть и конец этой встречи:
    // start меньше end по условию, поэтому endIdx остаётся допустимым.
    for startIdx < len(intervals) {
        // При равенстве сначала освобождаем зал, затем начинаем встречу.
        if start[startIdx] < end[endIdx] {
            startIdx++
            countRooms++
        } else {
            endIdx++
            countRooms--
        }
        // Пик активности — необходимое и достаточное число залов.
        if countRooms > maxRooms {
            maxRooms = countRooms
        }
    }

    return maxRooms
}
