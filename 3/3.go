package main

/*
Дана строка s, найдите длину самой длинной подстроки без повторяющихся символов.

Пример 1:
Input: s = "abcabcbb"
Output: 3
Пояснение: Ответ — "abc" длиной 3. Обратите внимание, что "bca" и "cab" также являются правильными ответами.

Пример 2:
Input: s = "bbbbb"
Output: 1
Пояснение: Ответ — "b" длиной 1.

Пример 3:
Input: s = "pwwkew"
Output: 3
Пояснение: Ответ: "wke" длиной 3.
Обратите внимание, что ответ должен быть подстрокой, "pwke" — это подпоследовательность, а не подстрока.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

*/

func lengthOfLongestSubstring(s string) int {
	maxCount := 0
    l, r := 0, 0
	hash := make(map[byte]struct{}, 0)
	for r < len(s) {
		if _, ok := hash[s[r]]; ok {
			delete(hash, s[l])
			l++
		} else {
			hash[s[r]] = struct{}{}
			maxCount = max(maxCount, r - l + 1)
			r++
		}
	}
	return maxCount
}
