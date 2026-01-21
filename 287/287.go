package main

/*
Дан массив целых чисел nums, содержащий n + 1 целых чисел, где каждое целое число находится в диапазоне [1, n] включительно.
В nums есть только одно повторяющееся число, верните это повторяющееся число.
Вы должны решить проблему, не изменяя числа массива и используя только постоянное дополнительное пространство.

Пример 1:
Input: nums = [1,3,4,2,2]
Output: 2

Пример 2:
Input: nums = [3,1,3,4,2]
Output: 3

Пример 3:
Input: nums = [3,3,3,3,3]
Output: 3

Constraints:
1 <= n <= 10^5
nums.length == n + 1
1 <= nums[i] <= n
Все целые числа в nums встречаются только один раз, за ​​исключением одного целого числа, которое встречается два или более раз.

*/

func findDuplicate(nums []int) int {
    // Floyd's algo
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0
	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}
	return slow2
}
