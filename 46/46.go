package main

/*
Дан массив различных целых чисел, верните все возможные перестановки.
Вы можете вернуть ответ в любом порядке.

Пример 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Пример 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Пример 3:
Input: nums = [1]
Output: [[1]]

Ограничения:
1 <= nums.length <= 6
-10 <= nums[i] <= 10
Все целые числа nums уникальны.

*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res := make([][]int, 0)
	perms := permute(nums[1:])
	for _, p := range perms {
		for i := range len(p)+1 {
			newPerm := make([]int, 0, len(p)+1)
            newPerm = append(newPerm, p[:i]...)
            newPerm = append(newPerm, nums[0])
            newPerm = append(newPerm, p[i:]...)
            res = append(res, newPerm)
		}
	}
	return res
}
