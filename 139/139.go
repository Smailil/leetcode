package main

/*
Дана строка s и словарь строк wordDict, верните true, если s можно сегментировать в разделенную пробелами последовательность из одного или нескольких словарных слов.
Обратите внимание, что одно и то же слово в словаре может использоваться при сегментации несколько раз.

Пример 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Объяснение: Возвращайте true, потому что "leetcode" можно сегментировать как "leet code".

Пример 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Объяснение: Возвращайте true, поскольку "applepenapple" можно сегментировать как "apple pen apple".
Обратите внимание, что вам разрешено повторно использовать словарное слово.

Пример 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false

Ограничения:
1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s и wordDict[i] состоят только из строчных английских букв.
Все строки wordDict уникальны.

*/

func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[len(s)] = true

    for i := len(s) - 1; i >= 0; i-- {
        for _, w := range wordDict {
            if i+len(w) <= len(s) && s[i:i+len(w)] == w {
                dp[i] = dp[i+len(w)]
            }
            if dp[i] {
                break
            }
        }
    }

    return dp[0]
}
