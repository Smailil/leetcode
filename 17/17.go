package main

/*
Дана строка, содержащая цифры от 2 до 9 включительно, верните все возможные комбинации букв, которые может представлять число. Верните ответ в любом порядке.
Соответствие цифр буквам (как на кнопках телефона) приведено ниже. Обратите внимание, что 1 не соответствует никаким буквам.
https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png

Пример 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Пример 2:
Input: digits = "2"
Output: ["a","b","c"]

Ограничения:
1 <= digits.length <= 4
digits[i] — цифра в диапазоне ['2', '9'].

*/

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
		return []string{}
	}
	var result []string
	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtrack func(int, string)
	backtrack = func(i int, curS string) {
		if len(digits) == i {
			result = append(result, curS)
			return 
		}
		letters := phoneMap[digits[i]-'2']
		for _, c := range letters {
			backtrack(i+1, curS+string(c))
		}
	}
	backtrack(0, "")
	return result
}
