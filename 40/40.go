package main

import "sort"

/*
Дан набор номеров кандидатов (кандидатов) и целевое число (цель), найдите все уникальные комбинации среди кандидатов,
в которых сумма чисел кандидатов равна целевому значению.
Каждое число кандидатов можно использовать в комбинации только один раз.
Примечание. Набор решений не должен содержать повторяющихся комбинаций.

Пример 1:
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Пример 2:
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Ограничения:
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

*/

func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(int, []int, int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			copyCur := make([]int, len(cur))
			copy(copyCur, cur)
			res = append(res, copyCur)
			return 
		}
		if total > target || i == len(candidates) {
			return
		}
		cur = append(cur, candidates[i])
		dfs(i + 1, cur, total+candidates[i])
		cur = cur[:len(cur)-1]
		for i + 1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, cur, total)
	}
	dfs(0, make([]int, 0), 0)
	return res
}
