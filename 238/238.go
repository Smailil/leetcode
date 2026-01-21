package main

/*
Дан целочисленный массив nums, верните массив answer таким образом, 
чтобы answer[i] был равен произведению всех элементов nums, кроме nums[i].

Произведение любого префикса или суффикса nums гарантированно укладывается в 32-bit integer.

Вы должны написать алгоритм, который выполняется за O(n) времени и без использования операции деления.

Пример 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Пример 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Ограничения:
2 <= nums.length <= 105
-30 <= nums[i] <= 30
Входные данные генерируются таким образом, что answer[i] гарантированно укладывается в 32-bit integer.

*/

func productExceptSelf(nums []int) []int {
    output := make([]int, len(nums))

	pref := 1
	for i, num := range nums {
		output[i] = pref
		pref *= num
	}
	suf := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= suf
		suf *= nums[i]
	}
	return output
}