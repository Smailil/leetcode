package main

/*
Существует целочисленный массив nums, отсортированный по возрастанию (с разными значениями).
Перед передачей в функцию nums, возможно, поворачивается влево по неизвестному индексу k (1 <= k < nums.length),
так что результирующий массив имеет вид [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (индексирован 0). 
Например, [0,1,2,4,5,6,7] можно повернуть влево на 3 индекса и превратить в [4,5,6,7,0,1,2].
Даны массив nums после возможного поворота и целочисленная цель, верните индекс цели, если он находится в nums, или -1, если он не в nums.
Вы должны написать алгоритм со сложностью выполнения O(log n).

Пример 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Пример 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Пример 3:
Input: nums = [1], target = 0
Output: -1

Ограничения:
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
Все значения чисел уникальны.
nums — это возрастающий массив, который возможно вращается.
-10^4 <= target <= 10^4

*/

func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return mid
		}
		// left
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		// right
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
