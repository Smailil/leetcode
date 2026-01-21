package main

import "sort"

/*
Дан целочисленный массив nums, который может содержать дубликаты, верните все возможные подмножества (набор мощности).
Набор решений не должен содержать повторяющихся подмножеств. Верните решение в любом порядке.

Пример 1:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Пример 2:
Input: nums = [0]
Output: [[],[0]]

Ограничения:
1 <= nums.length <= 10
-10 <= nums[i] <= 10

*/

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var backtrack func(int, []int)
	backtrack = func(i int, cur []int) {
		if i == len(nums) {
			copyCur := make([]int, len(cur))
			copy(copyCur, cur)
			res = append(res, copyCur)
			return
		}
		cur = append(cur, nums[i])
		backtrack(i+1, cur)
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i+1, cur)
	}
	backtrack(0, make([]int, 0))
	return res
}
