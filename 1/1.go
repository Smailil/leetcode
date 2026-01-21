package main

func twoSum(nums []int, target int) []int {
	viewedNums := make(map[int]int)
	for i, num := range nums {
		oldNum := target - num
		oldi, ok := viewedNums[oldNum]
		if ok {
			return []int{oldi, i}
		}
		viewedNums[num] = i
	}
	return []int{}
}
