package main

/*
Дана строка s, нужно вернуть самую длинную палиндромную подстроку в s.

Пример 1:
Input: s = "babad"
Output: "bab"
Объяснение: "aba" также является допустимым ответом.

Пример 2:
Input: s = "cbbd"
Output: "bb"

Ограничения:
1 <= s.length <= 1000
s состоят только из цифр и английских букв.

*/

func longestPalindrome(s string) string {
    resIdx, resLen := 0, 0

    for i := range len(s) {
        // odd length
        l, r := i, i
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r-l+1 > resLen {
                resIdx = l
                resLen = r - l + 1
            }
            l--
            r++
        }

        // even length
        l, r = i, i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            if r-l+1 > resLen {
                resIdx = l
                resLen = r - l + 1
            }
            l--
            r++
        }
    }

    return s[resIdx : resIdx+resLen]
}