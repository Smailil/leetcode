package main

/*
Дан массив целых чисел nums, отсортированный в порядке возрастания, и целочисленную цель, напишите функцию для поиска цели в nums. 
Если цель существует, верните ее индекс. В противном случае верните -1.
Вы должны написать алгоритм с O(log n)

Пример 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Объяснение: число 9 существует в nums, а его индекс равен 4.

Пример 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Объяснение: числа 2 не существует в nums, поэтому верните -1.

Ограничения:
1 <= nums.length <= 10^4
-10^4 < nums[i], target < 10^4
Все целые числа в nums уникальны.
nums отсортирован в порядке возрастания.

*/

func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
	for l <= r {
		mid := (l+r)/2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
