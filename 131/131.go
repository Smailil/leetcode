package main

/*
Дана строка s, разбиения s должны быть такие, что каждая подстрока раздела является палиндромом. Вернуть все возможные палиндромные разбиения s.

Пример 1:
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Пример 2:
Input: s = "a"
Output: [["a"]]

Ограничения,:
1 <= s.length <= 16
s содержит только строчные английские буквы.

*/

func isPal(s string) bool {
	i, j := 0, len(s)
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func partition(s string) [][]string {
    res := make([][]string, 0)
	part := make([]string, 0, len(s))
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(s) {
			cPart := make([]string, len(part))
			copy(cPart, part)
			res = append(res, cPart)
			return 
		}
		for j := i; j < len(s); j++ {
			if p := s[i:j+1] ; isPal(p) {
				part = append(part, p)
				dfs(j+1)
				part = part[:len(part)-1]
			}
		}
	}
	dfs(0)
	return res
}
