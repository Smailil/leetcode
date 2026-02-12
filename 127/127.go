package main

/*
Последовательность преобразования слова BeginWord в слово EndWord с использованием словаря wordList — это последовательность слов BeginWord -> s1 -> s2 -> ... -> sk такая, что:
Каждая соседняя пара слов отличается одной буквой.
Каждый si для 1 <= i <= k находится в списке слов. Обратите внимание, что BeginWord не обязательно должен находиться в списке слов.
sk == конечное слово
Учитывая два слова, BeginWord и EndWord, и словарь WordList, верните количество слов в кратчайшей последовательности преобразования из BeginWord в EndWord,
или 0, если такой последовательности не существует.

Пример 1:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Пояснение: Одна из самых коротких последовательностей преобразований — «hit» -> «hot» -> «dot» -> «dog» -> cog», длина которой составляет 5 слов.

Пример 2:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Объяснение: Конечное слово "cog" отсутствует в списке слов, поэтому не существует допустимой последовательности преобразования.

Ограничения:
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord и wordList[i] состоят из строчных английских букв.
beginWord != endWord
Все слова в wordList уникальны.

*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
    if !contains(wordList, endWord) {
        return 0
    }

    nei := make(map[string][]string)
    wordList = append(wordList, beginWord)

    for _, word := range wordList {
        for j := range len(word) {
            pattern := word[:j] + "*" + word[j+1:]
            nei[pattern] = append(nei[pattern], word)
        }
    }

    visit := make(map[string]struct{})
    visit[beginWord] = struct{}{}
    q := []string{beginWord}
    res := 1

    for len(q) > 0 {
        for i := len(q); i > 0; i-- {
            word := q[0]
            q = q[1:]

            if word == endWord {
                return res
            }

            for j := range len(word) {
                pattern := word[:j] + "*" + word[j+1:]
                for _, neiWord := range nei[pattern] {
                    if _, ok := visit[neiWord]; !ok {
                        visit[neiWord] = struct{}{}
                        q = append(q, neiWord)
                    }
                }
            }
        }
        res++
    }
    return 0
}

func contains(wordList []string, word string) bool {
    for _, w := range wordList {
        if w == word {
            return true
        }
    }
    return false
}
