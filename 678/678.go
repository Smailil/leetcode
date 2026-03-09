package main

/*
Дана строка s, содержащую только три типа символов: '(', ')' и '*', верните true, если s допустима.

Следующие правила определяют допустимую строку:
Любая левая скобка '(' должна иметь соответствующую правую скобку ')'.
Любая правая скобка ')' должна иметь соответствующую левую скобку '('.
Левая скобка '(' должна стоять перед соответствующей правой скобкой ')'.
'*' можно рассматривать как одну правую скобку ')' или одну левую скобку '(' или пустую строку "".

Пример 1:
Input: s = "()"
Output: true

Пример 2:
Input: s = "(*)"
Output: true

Пример 3:
Input: s = "(*))"
Output: true

Ограничения:
1 <= s.length <= 100
s[i] — это '(', ')' или '*'.

*/

func checkValidString(s string) bool {
    leftMin, leftMax := 0, 0

    for _, c := range s {
        switch c {
		case '(':
            leftMin, leftMax = leftMin+1, leftMax+1
        case ')':
            leftMin, leftMax = leftMin-1, leftMax-1
        default:
            leftMin, leftMax = leftMin-1, leftMax+1
        }
        if leftMax < 0 {
            return false
        }
        if leftMin < 0 {
            leftMin = 0
        }
    }
    return leftMin == 0
}
