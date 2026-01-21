package main

import "strings"

/*
Дан n пар круглых скобок, напишите функцию, которая генерирует все комбинации правильных круглых скобок.

Пример 1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Пример 2:
Input: n = 1
Output: ["()"]

Ограничения:
1 <= n <= 8

*/

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	stack := make([]string, 0, n*2)
    var backtrack func (int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			res = append(res, strings.Join(stack, ""))
			return 
		}
		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}
		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return res
}
