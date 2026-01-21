package main

/*
Учитывая целочисленный массив уникальных элементов, верните все возможные подмножества (набор мощности).
Набор решений не должен содержать повторяющихся подмножеств. Верните решение в любом порядке.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]
 
Ограничения:
1 <= nums.length <= 10
-10 <= nums[i] <= 10
Все числа nums уникальны.

*/

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }

    if m == 1 {
        return n
    }

    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}


func subsets(nums []int) [][]int {
    res := make([][]int, 0, IntPow(2, len(nums)))
	subset := make([]int, 0, len(nums))
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			copySubset := make([]int, len(subset))
			copy(copySubset, subset)
			res = append(res, copySubset)
			return
		}
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	
	}
	dfs(0)
	return res
}
