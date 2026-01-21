package main

import "fmt"

func longestSubarray(nums []int) int {
	prev, curr, maxSum := 0, 0, 0
	hasZero := false

	for _, num := range nums {
		if num == 1 {
			curr++
		} else {
			hasZero = true
			maxSum = max(maxSum, prev+curr)
			prev = curr
			curr = 0
		}
	}

	maxSum = max(maxSum, prev+curr)

	if hasZero {
		return maxSum
	}

	return maxSum - 1
}

func main() {
	a := longestSubarray([]int{1, 0})
	fmt.Println(a)
}
