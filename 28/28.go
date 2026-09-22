package main

/*
Даны две строки needle и haystack. Верните индекс первого вхождения needle
в haystack или -1, если needle не является подстрокой haystack.

Пример 1:
Вход: haystack = "sadbutsad", needle = "sad"
Выход: 0
Пояснение: "sad" встречается по индексам 0 и 6.
Первое вхождение находится по индексу 0, поэтому возвращаем 0.

Пример 2:
Вход: haystack = "leetcode", needle = "leeto"
Выход: -1
Пояснение: "leeto" не встречается в "leetcode", поэтому возвращаем -1.

Ограничения:
1 <= haystack.length, needle.length <= 10^4
haystack и needle состоят только из строчных английских букв.
*/

// strStr возвращает индекс первого вхождения needle в haystack. Двойной
// полиномиальный хеш позволяет сравнивать окна длины m за O(1).
// Время: O(n + m), где n = len(haystack), а m = len(needle).
// Дополнительная память: O(1).
func strStr(haystack string, needle string) int {
    // Пустая строка считается вхождением в начало. Ограничения задачи
    // исключают этот случай, но функция обрабатывает его явно.
    if needle == "" {
        return 0
    }

    // Две пары основания и модуля уменьшают вероятность коллизии:
    // совпадение обоих хешей используется как признак равенства строк.
    base1, mod1 := int64(31), int64(768258391)
    base2, mod2 := int64(37), int64(685683731)

    n, m := len(haystack), len(needle)
    // Если needle длиннее haystack, окно длины m построить невозможно.
    if m > n {
        return -1
    }

    // power хранит base^m по соответствующему модулю. Этот множитель
    // понадобится, чтобы удалить вклад первого символа при сдвиге окна.
    power1, power2 := int64(1), int64(1)
    for range m {
        power1 = (power1 * base1) % mod1
        power2 = (power2 * base2) % mod2
    }

    var needleHash1, needleHash2 int64
    var haystackHash1, haystackHash2 int64

    // Одновременно вычисляем хеши needle и первого окна haystack[0:m].
    // Проверка m <= n выше гарантирует допустимость индекса haystack[i].
    for i := range m {
        needleHash1 = (needleHash1*base1 + int64(needle[i])) % mod1
        needleHash2 = (needleHash2*base2 + int64(needle[i])) % mod2
        haystackHash1 = (haystackHash1*base1 + int64(haystack[i])) % mod1
        haystackHash2 = (haystackHash2*base2 + int64(haystack[i])) % mod2
    }

    // В начале итерации хеши haystack соответствуют окну haystack[i:i+m].
    // Правая граница включительна: начало последнего допустимого окна равно
    // n - m. Это также гарантирует одну проверку, когда n == m.
    for i := 0; i <= n-m; i++ {
        // Двойное совпадение резко снижает вероятность коллизии, но код
        // не выполняет дополнительную посимвольную проверку строк.
        if haystackHash1 == needleHash1 && haystackHash2 == needleHash2 {
            return i
        }

        if i+m < n {
            // Умножение на base сдвигает степени. Затем вычитается первый
            // символ старого окна и добавляется следующий символ haystack.
            // Условие выше гарантирует, что индекс i + m находится в строке.
            // Хеши и степени меньше 10^9, а код символа не превышает 255,
            // поэтому промежуточные произведения помещаются в int64.
            haystackHash1 = (haystackHash1*base1 - int64(haystack[i])*power1 + int64(haystack[i+m])) % mod1
            haystackHash2 = (haystackHash2*base2 - int64(haystack[i])*power2 + int64(haystack[i+m])) % mod2

            // В Go остаток от отрицательного числа отрицателен, поэтому
            // возвращаем хеш в диапазон [0, mod) после вычитания символа.
            if haystackHash1 < 0 {
                haystackHash1 += mod1
            }
            if haystackHash2 < 0 {
                haystackHash2 += mod2
            }
        }
    }

    return -1
}
