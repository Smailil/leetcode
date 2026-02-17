package main

/*
Дана строка s, верните количество палиндромных подстрок в ней.
Строка является палиндромом, если она читается как в прямом, так и в обратном направлении.
Подстрока — это непрерывная последовательность символов внутри строки.

Example 1:
Input: s = "abc"
Output: 3
Пояснение: Три палиндромные строки: «a», «b», «c».

Example 2:
Input: s = "aaa"
Output: 6
Пояснение: Шесть палиндромных строк: «а», «а», «а», «аа», «аа», «ааа».

Ограничения:
1 <= s.length <= 1000
s состоит из строчных английских букв.

*/

func countSubstrings(s string) int {
    res := 0
    for i := 0; i < len(s); i++ {
        res += countPals(s, i, i)
        res += countPals(s, i, i+1)
    }
    return res
}

func countPals(s string, l, r int) int {
    res := 0
    for l >= 0 && r < len(s) && s[l] == s[r] {
        res++
        l--
        r++
    }
    return res
}
