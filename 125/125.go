package main

/*
Фраза является палиндромом, если после преобразования всех прописных букв в строчные и удаления всех небуквенно-цифровых 
символов она читается одинаково и вперед, и назад. Буквенно-цифровые символы включают буквы и цифры.

Дана строка s, верните true, если это палиндром, или false в противном случае.

Пример 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Объяснение: "amanaplanacanalpanama" — палиндром.

Пример 2:
Input: s = "race a car"
Output: false
Объяснение: "raceacar" — не палиндром.

Пример 3:
Input: s = " "
Output: true
Объяснение: s — пустая строка "" после удаления небуквенно-цифровых символов.
Поскольку пустая строка читается одинаково и вперед, и назад, она является палиндромом.

Ограничения:
1 <= s.length <= 2 * 105
s состоит только из печатных символов ASCII.

*/

func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func isLetter(b byte) bool {
	return ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9')
}

func isPalindrome(s string) bool {
    l := 0
	r := len(s)-1
	for l <= r {
		if !isLetter(s[l]) {
			l++
			continue
		}
		if !isLetter(s[r]) {
			r--
			continue
		}
		if toLower(s[l]) == toLower(s[r]) {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}