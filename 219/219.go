package main

/*
Дан массив целых чисел nums и целое число k.
Верните true, если существуют два различных индекса i и j такие,
что nums[i] == nums[j] и abs(i - j) <= k.

Пример 1:
Вход: nums = [1,2,3,1], k = 3
Выход: true

Пример 2:
Вход: nums = [1,0,1,1], k = 1
Выход: true

Пример 3:
Вход: nums = [1,2,3,1,2,3], k = 2
Выход: false

Ограничения:
1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5

*/

func containsNearbyDuplicate(nums []int, k int) bool {
    window := make(map[int]bool)
    L := 0

    for R := range nums {
        if R-L > k {
            delete(window, nums[L])
            L++
        }
        if window[nums[R]] {
            return true
        }
        window[nums[R]] = true
    }

    return false
}
