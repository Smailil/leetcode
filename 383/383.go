package main

/*
Даны две строки ransomNote и magazine.
Верните true, если строку ransomNote можно составить из букв строки magazine,
иначе верните false.

Каждую букву строки magazine можно использовать в ransomNote только один раз.

Пример 1:
Вход: ransomNote = "a", magazine = "b"
Выход: false

Пример 2:
Вход: ransomNote = "aa", magazine = "ab"
Выход: false

Пример 3:
Вход: ransomNote = "aa", magazine = "aab"
Выход: true

Ограничения:
1 <= ransomNote.length, magazine.length <= 10^5
ransomNote и magazine состоят из строчных английских букв.
*/

// canConstruct проверяет, хватает ли букв magazine для составления ransomNote.
// Алгоритм сначала сохраняет количество доступных букв, а затем расходует их.
// Время: O(m + n), где m = len(ransomNote), n = len(magazine).
// Дополнительная память: O(1), поскольку массив всегда содержит 26 счётчиков.
func canConstruct(ransomNote string, magazine string) bool {
    // Индекс каждой буквы задаётся её смещением от 'a', поэтому один массив
    // хранит доступное количество всех допустимых по условию символов.
    count := make([]int, 26)
    for _, c := range magazine {
        count[c-'a']++
    }

    // Для каждой требуемой буквы уменьшаем её запас. После обработки префикса
    // count хранит остаток букв magazine, не использованных этим префиксом.
    for _, c := range ransomNote {
        count[c-'a']--

        // Отрицательное значение означает, что текущая буква понадобилась чаще,
        // чем встречалась в magazine. Дальнейший обход уже ничего не изменит.
        if count[c-'a'] < 0 {
            return false
        }
    }

    // Все буквы успешно списаны, следовательно, ransomNote можно составить.
    return true
}