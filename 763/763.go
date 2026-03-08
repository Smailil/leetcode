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

    lastIndex := make([]int, 26)
    for i := 0; i < len(s); i++ {
        lastIndex[s[i]-'a'] = i
    }

    res := make([]int, 0)
    currentLen, partitionEnd := 0, 0

    for i := 0; i < len(s); i++ {
        currentLen++
        
        if lastIndex[s[i]-'a'] > partitionEnd {
            partitionEnd = lastIndex[s[i]-'a']
        }

        if partitionEnd == len(s)-1 {
            res = append(res, currentLen + (partitionEnd - i))
            break
        }

        if i == partitionEnd {
            res = append(res, currentLen)
            currentLen = 0
        }
    }

    return res
}
