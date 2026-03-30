package main

/*
Даны две строки s и t, верните количество различных подпоследовательностей s, равное t.
Тестовые примеры генерируются таким образом, чтобы ответ соответствовал 32-битному целому числу со знаком.

Пример 1:
Input: s = "rabbbit", t = "rabbit"
Output: 3
Объяснение:
Как показано ниже, существует три способа создания «rabbit» из s.
(rabb)b(it)
(ra)b(bbit)
(rab)b(bit)

Пример 2:
Input: s = "babgbag", t = "bag"
Output: 5
Объяснение:
Как показано ниже, существует 5 способов создания «bag» из s.
(ba)b(g)bag
(ba)bgba(g)
(b)abgb(ag)
ba(b)gb(ag)
babg(bag)

Ограничения:
1 <= s.length, t.length <= 1000
s и t состоят из английских букв.

*/

func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
    dp := make([]int, n+1)
    dp[n] = 1

    for i := m - 1; i >= 0; i-- {
        prev := dp[n]
        for j := n - 1; j >= 0; j-- {
            temp := dp[j]
            if s[i] == t[j] {
                dp[j] += prev
            }
            prev = temp
        }
    }

    return dp[0]
}