package main

/*
Вам даны n шариков, пронумерованных от 0 до n - 1. На каждом шарике нарисован номер, представленный массивом nums. Вас просят лопнуть все шарики.
Если вы лопнете i-й шарик, вы получите nums[i - 1] * nums[i] * nums[i + 1] монет. 
Если i - 1 или i + 1 выходит за пределы массива, относитесь к нему так, как будто на нем нарисован воздушный шар с цифрой 1.
Верните как можно больше монет, которые сможете собрать, разумно лопая воздушные шары.

Пример 1:
Input: nums = [3,1,5,8]
Output: 167
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins = 3*1*5 + 3*5*8 + 1*3*8 + 1*8*1 = 167

Пример 2:
Input: nums = [1,5]
Output: 10

Ограничения:
n == nums.length
1 <= n <= 300
0 <= nums[i] <= 100

*/

func maxCoins(nums []int) int {
    n := len(nums)
    arr := make([]int, n+2)
    arr[0] = 1
    arr[n+1] = 1
    copy(arr[1:], nums)

	n = len(arr)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    var dfs func(l, r int) int
    dfs = func(l, r int) int {
        if l > r {
            return 0
        }
        if dp[l][r] > 0 {
            return dp[l][r]
        }

        for i := l; i <= r; i++ {
            coins := arr[l-1] * arr[i] * arr[r+1]
            coins += dfs(l, i-1) + dfs(i+1, r)
            if dp[l][r] < coins {
                dp[l][r] = coins
            }
        }
        return dp[l][r]
    }

    return dfs(1, n-2)
}