package main

/*
Дан несортированный массив целых чисел nums, верните длину самой длинной последовательности последовательных элементов.

Вы должны написать алгоритм, который выполняется за O(n) времени.

Пример 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Объяснение: Самая длинная последовательная последовательность элементов — [1, 2, 3, 4]. Следовательно, его длина равна 4.

Пример 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9

Пример 3:
Input: nums = [1,0,1,2]
Output: 3

Ограничения:
0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/
func longestConsecutive(nums []int) int {
	maxLen := 0
    hm := make(map[int]struct{})
	for _, num := range nums {
		hm[num] = struct{}{}
	}
	for num := range hm {
		if _, ok := hm[num-1]; ok {
			continue
		}
		conLen := 1
		for {
			if _, ok := hm[num+conLen]; !ok {
				break
			}
			conLen++
		}
		maxLen = max(maxLen, conLen)
	}
	return maxLen
}