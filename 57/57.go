package main

/*
Вам дан массив непересекающихся интервалов, где интервалы [i] = [starti, endi] представляют начало и конец i-го интервала, 
а интервалы сортируются в порядке возрастания по starti. Вам также дан интервал newInterval = [start, end], который представляет начало и конец другого интервала.
Вставьте newInterval в интервалы таким образом, чтобы интервалы по-прежнему сортировались в порядке возрастания по начальным значениям, 
а интервалы по-прежнему не имели перекрывающихся интервалов (при необходимости объединяйте перекрывающиеся интервалы).
Верните интервалы после вставки.
Обратите внимание, что вам не нужно изменять интервалы на месте. Вы можете создать новый массив и вернуть его.

Пример 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Пример 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Пояснение: Потому что новый интервал [4,8] перекрывается с [3,5],[6,7],[8,10].

Ограничения:
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= starti <= endi <= 10^5
интервалы сортируются по starti в порядке возрастания.
newInterval.length == 2
0 <= start <= end <= 10^5

*/

func insert(intervals [][]int, newInterval []int) [][]int {
    res := make([][]int, 0, len(intervals)+1)

    for i := range intervals {
        if newInterval[1] < intervals[i][0] {
            res = append(res, newInterval)
            return append(res, intervals[i:]...)
        } else if newInterval[0] > intervals[i][1] {
            res = append(res, intervals[i])
        } else {
            // newInterval = []int{
            //     min(newInterval[0], intervals[i][0]),
            //     max(newInterval[1], intervals[i][1]),
            // }
			if intervals[i][0] < newInterval[0] {
			    newInterval[0] = intervals[i][0]
			}
			if intervals[i][1] > newInterval[1] {
    			newInterval[1] = intervals[i][1]
			}
        }
    }
    res = append(res, newInterval)
    return res
}