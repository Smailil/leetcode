package main

/*
Вам дан целочисленный массив coin, представляющий монеты разного номинала, и целочисленная сумма, представляющая общую сумму денег.
Верните наименьшее количество монет, которое вам нужно, чтобы составить эту сумму. Если эта сумма денег не может быть составлена ​​ни одной комбинацией монет, верните -1.
Вы можете предположить, что у вас есть бесконечное количество монет каждого вида.

Пример 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Пояснение: 11 = 5 + 5 + 1

Пример 2:
Input: coins = [2], amount = 3
Output: -1

Пример 3:
Input: coins = [1], amount = 0
Output: 0

Ограничения:
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4

*/

func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1 // max value
    }
    dp[0] = 0
    for a := 1; a <= amount; a++ {
        for _, c := range coins {
            if a-c >= 0 {
                dp[a] = min(dp[a], 1+dp[a-c])
            }
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}