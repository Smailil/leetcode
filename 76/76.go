package main

import "math"

/*
Даны две строки s и t длиной m и n соответственно, верните минимальную подстроку окна s, такую,
что каждый символ в t (включая дубликаты) включен в окно. Если такой подстроки нет, верните пустую строку «».
Тестовые примеры будут сгенерированы таким образом, чтобы ответ был уникальным.

Пример 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Объяснение: Подстрока минимального окна "BANC" включает "A", "B" и "C" из строки t.

Пример 2:
Input: s = "a", t = "a"
Output: "a"
Объяснение: Вся строка s представляет собой минимальное окно.

Пример 3:
Input: s = "a", t = "aa"
Output: ""
Объяснение: В окно должны быть включены обе буквы 'a' из t.
Поскольку самое большое окно s имеет только одну букву «а», верните пустую строку.

Ограничения:
m == s.length
n == t.length
1 <= m, n <= 10^5
s и t состоят из прописных и строчных английских букв.

*/
const GREATER = math.MaxInt32

func minWindow(s string, t string) string {
    if len(t) < 1 || len(s) < len(t) {
		return ""
	}
	countT := make(map[byte]int)
	for i:=0;i<len(t);i++ {
		countT[t[i]]++
	}
	window := make(map[byte]int)
	have, need := 0, len(countT)
	res, resLen := [2]int{-1, -1}, GREATER
	l, r := 0, 0
	for ;r < len(s); r++ {
		c := s[r]
		window[c]++

		if count, ok := countT[c]; ok && count == window[c]{
			have++
		}
		for have == need {
			if resLen > r - l + 1 {
				res = [2]int{l, r}
				resLen = r - l + 1
			}
			window[s[l]]--
			if count, ok := countT[s[l]]; ok && count - 1 == window[s[l]] {
				have--
			}
			l++
		}
	}
	if resLen == GREATER {
		return ""
	}
	return s[res[0]:res[1] + 1]
}
