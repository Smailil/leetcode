package main

import "strings"

/*
Даны две строки word1 и word2.
Объедините строки, поочерёдно добавляя символы,
начиная с word1.
Если одна строка длиннее другой, добавьте оставшиеся символы
в конец результирующей строки.

Верните объединённую строку.

Пример 1:
Вход: word1 = "abc", word2 = "pqr"
Выход: "apbqcr"
Пояснение: объединённая строка формируется следующим образом:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r

Пример 2:
Вход: word1 = "ab", word2 = "pqrs"
Выход: "apbqrs"
Пояснение: так как word2 длиннее, "rs" добавляется в конец.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s

Пример 3:
Вход: word1 = "abcd", word2 = "pq"
Выход: "apbqcd"
Пояснение: так как word1 длиннее, "cd" добавляется в конец.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d

Ограничения:
1 <= word1.length, word2.length <= 100
word1 и word2 состоят из строчных букв английского алфавита.
*/

// mergeAlternately поочерёдно пишет байты двух строк и дописывает остаток.
// По условию это ASCII-буквы, поэтому байт совпадает с одним символом.
// Амортизированное время O(m+n), m=len(word1), n=len(word2).
// Рабочая память при росте Builder O(m+n); результат O(m+n) в его буфере.
// String не копирует финальный буфер: его считают один раз, вход не меняется.
func mergeAlternately(word1 string, word2 string) string {
    var res strings.Builder
    i, j := 0, 0

    // Оба индекса проверены до чтения; каждый шаг добавляет ровно одну пару.
    for i < len(word1) && j < len(word2) {
        res.WriteByte(word1[i])
        res.WriteByte(word2[j])
        i++
        j++
    }

    // Цикл остановился на конце хотя бы одной строки: один из хвостов пуст.
    // i и j могут равняться длине, что допустимо для среза пустого хвоста.
    res.WriteString(word1[i:])
    res.WriteString(word2[j:])

    return res.String()
}
