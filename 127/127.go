package main

/*
Последовательность преобразования слова beginWord в слово endWord
с использованием словаря wordList — это последовательность слов
beginWord -> s_1 -> s_2 -> ... -> s_k такая, что:
Каждая соседняя пара слов отличается одной буквой.
Каждое s_i для 1 <= i <= k находится в wordList.
Обратите внимание, что beginWord не обязательно должен находиться в wordList.
s_k == endWord
По словам beginWord и endWord и словарю wordList верните количество слов
в кратчайшей последовательности преобразования из beginWord в endWord,
или 0, если такой последовательности не существует.

Пример 1:
Вход: beginWord = "hit", endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]
Выход: 5
Пояснение: одна из кратчайших последовательностей преобразования —
«hit» -> «hot» -> «dot» -> «dog» -> «cog», её длина составляет 5 слов.

Пример 2:
Вход: beginWord = "hit", endWord = "cog",
wordList = ["hot","dot","dog","lot","log"]
Выход: 0
Объяснение: конечное слово "cog" отсутствует в wordList,
поэтому допустимой последовательности преобразования не существует.

Ограничения:
1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord и wordList[i] состоят из строчных английских букв.
beginWord != endWord
Все слова в wordList уникальны.
*/

// ladderLength ищет кратчайший путь BFS по шаблонам с одной заменой.
// N = len(wordList)+1, L = len(beginWord): среднее время O(N*L^2).
// Создание и хеширование каждого из L шаблонов стоит O(L). В одном шаблоне
// не более 26 разных слов, плюс возможный дубль beginWord: обходы соседей
// не добавляют множитель N. Оценка поиска map средняя, не худшая гарантия.
// Рабочая память O(N*L^2) байт для ключей шаблонов и O(N*L) ссылок;
// ответ int занимает O(1). Исходные элементы wordList не меняются,
// но append может записать beginWord в общий буфер за длиной входного среза.
func ladderLength(beginWord string, endWord string, wordList []string) int {
    // Конечное слово должно входить в словарь, иначе допустимого пути нет.
    if !contains(wordList, endWord) {
        return 0
    }

    nei := make(map[string][]string)
    wordList = append(wordList, beginWord)

    // Слова с одним и тем же шаблоном отличаются только в позиции '*'.
    // j — индекс ASCII-байта: по условию все буквы английские и длины равны.
    // 0 <= j < L, поэтому срезы до j и от j+1 допустимы.
    for _, word := range wordList {
        for j := range len(word) {
            pattern := word[:j] + "*" + word[j+1:]
            nei[pattern] = append(nei[pattern], word)
        }
    }

    // visit отмечает слово при постановке в очередь: повторной обработки нет.
    // res считает слова в пути, включая beginWord, а не число рёбер.
    visit := make(map[string]struct{})
    visit[beginWord] = struct{}{}
    q := []string{beginWord}
    res := 1

    for len(q) > 0 {
        // Фиксируем размер слоя до добавления соседей. Новые слова остаются
        // для следующего слоя, поэтому первое endWord даёт кратчайший путь.
        for i := len(q); i > 0; i-- {
            word := q[0]
            q = q[1:]

            if word == endWord {
                return res
            }

            // Перебираем все шаблоны; бакеты не очищаются, но каждый из них
            // может читаться лишь по числу слов этого бакета (не более 27).
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

// contains линейно проверяет точное равенство слова словарному элементу.
// N = len(wordList), L — длина слов: время O(N*L) в худшем случае
// из-за сравнения строк, дополнительная память O(1). Входы не меняются.
func contains(wordList []string, word string) bool {
    for _, w := range wordList {
        if w == word {
            return true
        }
    }
    return false
}
