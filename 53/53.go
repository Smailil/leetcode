package main

/*
Дан целочисленный массив nums, найдите подмассив с наибольшей суммой и верните его сумму.

Пример 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Пояснение: Подмассив [4,-1,2,1] имеет наибольшую сумму 6.

Пример 2:
Input: nums = [1]
Output: 1
Пояснение: Подмассив [1] имеет наибольшую сумму 1.

Пример 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Пояснение: Подмассив [5,4,-1,7,8] имеет наибольшую сумму 23.

Ограничения:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

*/

func maxSubArray(nums []int) int {
    maxSub, curSum := nums[0], 0
    for _, num := range nums {
        if curSum < 0 {
            curSum = 0
        }
        curSum += num
		if curSum > maxSub {
			maxSub = curSum
		}
    }
    return maxSub
}