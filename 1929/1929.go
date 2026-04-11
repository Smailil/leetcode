package main

/*
Дан целочисленный массив nums длиной n.
Необходимо создать массив ans длиной 2n,
где ans[i] == nums[i] и ans[i + n] == nums[i] для 0 <= i < n (индексация с нуля).
Иными словами, ans — это конкатенация двух массивов nums.
Верните массив ans.

Пример 1:
Вход: nums = [1,2,1]
Выход: [1,2,1,1,2,1]
Пояснение: Массив ans формируется следующим образом:
- ans = [nums[0],nums[1],nums[2],nums[0],nums[1],nums[2]]
- ans = [1,2,1,1,2,1]

Пример 2:
Вход: nums = [1,3,2,1]
Выход: [1,3,2,1,1,3,2,1]
Пояснение: Массив ans формируется следующим образом:
- ans = [nums[0],nums[1],nums[2],nums[3],nums[0],nums[1],nums[2],nums[3]]
- ans = [1,3,2,1,1,3,2,1]

Ограничения:
n == nums.length
1 <= n <= 1000
1 <= nums[i] <= 1000

*/

func getConcatenation(nums []int) []int {
    n := len(nums)
    ans := make([]int, 2*n)
    for i, num := range nums {
        ans[i] = num
        ans[i+n] = num
    }
    return ans
}
