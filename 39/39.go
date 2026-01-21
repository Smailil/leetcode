package main

/*
Дан массив различных целых чисел-кандидатов и целевое целочисленное значение, верните список всех уникальных комбинаций кандидатов, 
в которых сумма выбранных чисел равна целевому значению. Вы можете возвращать комбинации в любом порядке.
Одно и то же число может быть выбрано из кандидатов неограниченное количество раз. Две комбинации уникальны, если частота хотя бы одного из выбранных чисел различна.
Тестовые примеры генерируются таким образом, чтобы количество уникальных комбинаций, дающих целевое значение, составляло менее 150 комбинаций для данного входного сигнала.
 
Пример 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Объяснение:
2 и 3 являются кандидатами, а 2 + 2 + 3 = 7. Обратите внимание, что 2 можно использовать несколько раз.
7 — кандидат, а 7 = 7.
Это единственные две комбинации.

Пример 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Пример 3:
Input: candidates = [2], target = 1
Output: []
 
Ограничения:
1 <= candidates.length <= 30
2 <= candidates[i] <= 40
Все элементы кандидатов различны.
1 <= target <= 40

*/

func combinationSum(candidates []int, target int) [][]int {
    res := make([][]int, 0)
	var dfs func(int, []int, int)
	dfs = func(i int, cur []int, total int) {
		if total == target {
			r := make([]int, len(cur))
			copy(r, cur)
			res = append(res, r)
			return 
		}
		if i >= len(candidates) || total > target {
			return 
		}

		cur = append(cur, candidates[i])
		dfs(i, cur, total + candidates[i])
		cur = cur[:len(cur)-1]
		dfs(i + 1, cur, total)
	}
	dfs(0, make([]int, 0), 0)
	return res
}
