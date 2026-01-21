package main

import "strconv"

/*
Вам дан массив строковых токенов, который представляет арифметическое выражение в обратной польской нотации.
Вычислите выражение. Верните целое число, представляющее значение выражения.

Обратите внимание, что:
Допустимыми операторами являются '+', '-', '*', и '/'.
Каждый операнд может быть целым числом или другим выражением.
Деление между двумя целыми числами всегда округляется вниз.
Никакого деления на ноль не будет.
Входные данные представляют собой допустимое арифметическое выражение в обратной польской системе счисления.
Ответ и все промежуточные вычисления могут быть представлены в виде 32-разрядного целого числа.

Пример 1:
Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9

Пример 2:
Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6

Пример 3:
Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22

Ограничения:
1 <= tokens.length <= 104
tokens[i] — это либо оператор: «+», «-», «*» или «/», либо целое число в диапазоне [-200, 200] в виде строки.

*/

func pop(st *[]int) int {
	num := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return num
}

func push(st *[]int, num int) {
	*st = append(*st, num)
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, t := range tokens {
		switch t {
		case "+":
			num2 := pop(&stack)
			num1 := pop(&stack)
			push(&stack, num1 + num2)
		case "-":
			num2 := pop(&stack)
			num1 := pop(&stack)
			push(&stack, num1 - num2)
		case "*":
			num2 := pop(&stack)
			num1 := pop(&stack)
			push(&stack, num1 * num2)
		case "/":
			num2 := pop(&stack)
			num1 := pop(&stack)
			push(&stack, num1 / num2)
		default:
			num, _ := strconv.Atoi(t)
			push(&stack, num)
		}
	}
	return pop(&stack)

}