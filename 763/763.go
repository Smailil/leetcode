package main

/*
Вам дана строка s. Мы хотим разделить строку на как можно большее количество частей, чтобы каждая буква появлялась не более чем в одной части.
Например, строка "ababcc" может быть разделена на ["abab", "cc"], но такие разделы, как ["aba", "bcc"] или ["ab", "ab", "cc"] недействительны.
Обратите внимание, что разбиение выполнено так, что после объединения всех частей по порядку результирующая строка должна быть s.
Верните список целых чисел, представляющих размер этих частей.

Пример 1:
Input: s = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Объяснение:
Раздел — «ababcbaca», «defegde», «hijhklij».
Это раздел, в котором каждая буква встречается не более чем в одной части.
Раздел типа «ababcbacadefegde», «hijhklij» неверен, поскольку он разбивает s на меньшее количество частей.

Пример 2:
Input: s = "eccbbbbdec"
Output: [10]

Ограничения:
1 <= s.length <= 500
s состоит из строчных английских букв.

*/

func partitionLabels(s string) []int {
    if len(s) == 0 {
        return []int{}
    }

    // 1. Находим последнее вхождение каждого символа
    lastIndex := make([]int, 26) // Оптимизация: array вместо map для lowercase letters
    for i := range len(s) {
        lastIndex[s[i]-'a'] = i
    }

    // 2. Предварительно выделяем память для результата
    res := make([]int, 0, len(s)/2+1)
    currentLen, partitionEnd := 0, 0

    // 3. Проходим по строке и определяем границы разделов
    for i := range len(s) {
        currentLen++
        charIdx := s[i] - 'a'
        
        // Расширяем границу раздела до последнего вхождения текущего символа
        if lastIndex[charIdx] > partitionEnd {
            partitionEnd = lastIndex[charIdx]
        }

        // Достигли конца текущего раздела
        if i == partitionEnd {
            res = append(res, currentLen)
            currentLen = 0
        }
    }

    return res
}