package main

/*
Даны числа целочисленного массива, верните true, если вы можете разделить массив на два подмножества так, 
чтобы сумма элементов в обоих подмножествах была равна или false в противном случае.

Пример 1:
Input: nums = [1,5,11,5]
Output: true
Объяснение: Массив можно разделить на [1, 5, 5] и [11].

Пример 2:
Input: nums = [1,2,3,5]
Output: false
Объяснение: Массив нельзя разделить на подмножества с одинаковой суммой.

Ограничения:
1 <= nums.length <= 200
1 <= nums[i] <= 100

*/

// func canPartition(nums []int) bool {
//     total := 0
//     for _, num := range nums {
//         total += num
//     }
//     if total%2 != 0 {
//         return false
//     }

//     target := total / 2
// 	dp := map[int]struct{}{0: {}}

//     for i := len(nums) - 1; i >= 0; i-- {
//         nextDP := map[int]struct{}{}
//         for t := range dp {
//             if t+nums[i] == target {
//                 return true
//             } else if t+nums[i] < target {
// 				nextDP[t+nums[i]] = struct{}{}
// 			}
//             nextDP[t] = struct{}{}
//         }
//         dp = nextDP
//     }

//     return false
// }

func canPartition(nums []int) bool {
    total := 0
    for _, num := range nums {
        total += num
    }
    if total%2 != 0 {
        return false
    }

    target := total / 2
    dp := make([]bool, target+1)
    dp[0] = true

    for _, num := range nums {
        for j := target; j >= num; j-- {
            if dp[j-num] {
                dp[j] = true
                if j == target {
                    return true
                }
            }
        }
    }

    return dp[target]
}