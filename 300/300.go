package main

import "sort"

/*
Дан целочисленный массив nums, верните длину самой длинной строго возрастающей подпоследовательности.

Пример 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Пояснение: Самая длинная возрастающая подпоследовательность — [2,3,7,101], поэтому длина равна 4.

Пример 2:

Input: nums = [0,1,0,3,2,3]
Output: 4

Пример 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

Ограничения:
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4

*/

// func lengthOfLIS(nums []int) int { // O(n^2)
//     LIS := make([]int, len(nums))
//     maxLen := 1

//     for i := range LIS {
//         LIS[i] = 1
//     }

//     for i := len(nums) - 1; i >= 0; i-- {
//         for j := i + 1; j < len(nums); j++ {
//             if nums[i] < nums[j] {
// 				LIS[i] = max(LIS[i], LIS[j] + 1)
//             }
//         }
// 		maxLen = max(maxLen, LIS[i])
//     }

//     return maxLen
// }

func lengthOfLIS(nums []int) int { // O(n*log(n))
    
    // tails[i] = минимальный последний элемент 
    // для возрастающей подпоследовательности длины i+1
    tails := []int{}
    
    for _, num := range nums {
        // Бинарный поиск: ищем первый элемент >= num
        // sort.SearchInts возвращает индекс для вставки
        pos := sort.SearchInts(tails, num)
        
        if pos == len(tails) {
            // num больше всех элементов → увеличиваем длину LIS
            tails = append(tails, num)
        } else {
            // Нашли элемент ≥ num → заменяем его на num
            // Это даёт нам "более выгодный" последний элемент
            // для подпоследовательности той же длины
            tails[pos] = num
        }
    }
    
    // Длина tails = длина самой длинной возрастающей подпоследовательности
    return len(tails)
}