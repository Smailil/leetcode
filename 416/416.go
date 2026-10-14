package main

/*
Дан целочисленный массив nums.
Верните true, если массив можно разделить на два подмножества
с равными суммами элементов, иначе верните false.

Пример 1:
Вход: nums = [1,5,11,5]
Выход: true
Объяснение: Массив можно разделить на [1, 5, 5] и [11].

Пример 2:
Вход: nums = [1,2,3,5]
Выход: false
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

// canPartition ищет подмножество с суммой, равной половине общей суммы.
// Булева динамика обрабатывает каждый элемент ровно один раз.
// Время O(n*S+n), n=len(nums), S — половина общей суммы; рабочая память O(S).
// Результат bool O(1); nums не меняется, сумма до 20000 безопасна для int32.
func canPartition(nums []int) bool {
    total := 0
    for _, num := range nums {
        total += num
    }
    // Две целые одинаковые суммы невозможны при нечётном total.
    if total%2 != 0 {
        return false
    }

    target := total / 2
    dp := make([]bool, target+1)
    // Нулевая сумма достижима пустым подмножеством.
    dp[0] = true

    for _, num := range nums {
        // Идём вниз: dp[j-num] ещё относится к прошлому набору элементов.
        // Поэтому текущий num нельзя использовать дважды; j-num неотрицателен.
        for j := target; j >= num; j-- {
            if dp[j-num] {
                dp[j] = true
                // Найденная половина достаточна: сумма всех остальных тоже S,
                // даже если оставшиеся элементы ещё не прошли динамику.
                if j == target {
                    return true
                }
            }
        }
    }

    return dp[target]
}