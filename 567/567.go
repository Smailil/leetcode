package main

/*
Даны две строки s1 и s2, верните true, если s2 содержит перестановку s1, или false в противном случае.
Другими словами, верните true, если одна из перестановок s1 является подстрокой s2.

Пример 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Пояснение: s2 содержит одну перестановку s1 ("ba").

Пример 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false

Ограничения:
1 <= s1.length, s2.length <= 104
s1 и s2 состоят из строчных английских букв.

*/

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
		return false
	}
	countS1, countS2 := make([]int, 'z'-'a'+1), make([]int, 'z'-'a'+1)
	l, r := 0, 0
	for ; r < len(s1); r++ {
		countS1[s1[r]-'a']++
		countS2[s2[r]-'a']++
	}
	match := 0
	for i, count := range countS1 {
		if count == countS2[i] {
			match++
		}
	}
	for ;r < len(s2); r++ {
		if match == 'z'-'a'+1 {
			return true
		}
		i := s2[r]-'a'
		countS2[i]++
		if countS1[i] == countS2[i] {
			match++
		} else if countS1[i] + 1 == countS2[i] {
			match--
		}
		i = s2[l]-'a'
		countS2[i]--
		if countS1[i] == countS2[i] {
			match++
		} else if countS1[i] - 1 == countS2[i] {
			match--
		}
		l++

	}
	return match == 'z'-'a'+1
}
