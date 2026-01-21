package main
/*
Дан целочисленный массив nums и целое число k, верните k наиболее часто встречающихся элементов. 
Вы можете вернуть ответ в любом порядке.

Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:
Input: nums = [1], k = 1
Output: [1]

*/
func topKFrequent(nums []int, k int) []int {
    hm := make(map[int]int)
	for _, num := range nums {
		hm[num]++
	}
	buckets := make([][]int, len(nums) + 1)
	for num, count := range hm {
		buckets[count] = append(buckets[count], num)
	}
	resultArray := make([]int, 0, k)
	for i := len(nums); i != 0; i-- {
		for _, num := range buckets[i] {
			resultArray = append(resultArray, num)
			if len(resultArray) == k {
				return resultArray
			}
		}
	}
	return resultArray
}