package main

import "strings"

/*
Дана строка s, содержащая только символы '(', ')', '{', '}', '[' и ']'. Определите, является ли входная строка допустимой.
Входная строка допустима, если:
Открытые скобки закрываются скобками того же типа.
Открытые скобки закрываются в правильном порядке.
Каждой закрывающей скобке соответствует открытая скобка того же типа.

Пример 1:
Input: s = "()"
Output: true

Пример 2:
Input: s = "()[]{}"
Output: true

Пример 3:
Input: s = "(]"
Output: false

Пример 4:
Input: s = "([])"
Output: true

Пример 5:
Input: s = "([)]"
Output: false

Ограничения:
1 <= s.length <= 104
s состоит только из скобок '()[]{}'.

*/

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		if open, isClose := hash[c]; isClose {
			lenStack := len(stack)
			if lenStack > 0 && stack[lenStack-1] == open {
				stack = stack[:lenStack-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
