package main

/*
Напишите функцию для нахождения наидлиннейшего общего префикса
среди массива строк.

Если общего префикса нет, верните пустую строку "".

Пример 1:
Вход: strs = ["flower","flow","flight"]
Выход: "fl"

Пример 2:
Вход: strs = ["dog","racecar","car"]
Выход: ""
Пояснение: Среди входных строк нет общего префикса.

Ограничения:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] состоит только из строчных букв английского алфавита, если непуста.

*/

func longestCommonPrefix(strs []string) string {
    for i := range len(strs[0]) {
        for _, s := range strs {
            if i == len(s) || s[i] != strs[0][i] {
                return s[:i]
            }
        }
    }
    return strs[0]
}
