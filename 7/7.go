package main

import "math"

/*
Дано 32-битное целое число x со знаком, верните x с обратными цифрами.
Если изменение x приводит к выходу значения за пределы диапазона 32-битных целых чисел со знаком [-2^31, 2^31 - 1], верните 0.
Предположим, что среда не позволяет хранить 64-битные целые числа (со знаком или без знака).

Пример 1:
Input: x = 123
Output: 321

Пример 2:
Input: x = -123
Output: -321

Пример 3:
Input: x = 120
Output: 21

Ограничения:
-2^31 <= x <= 2^31 - 1

*/

func reverse(x int) int {
    MIN := -2147483648 // -2^31
    MAX := 2147483647  // 2^31 - 1

    res := 0
    for x != 0 {
        digit := int(math.Mod(float64(x), 10))
        x = int(float64(x) / 10)

        if res > MAX/10 || (res == MAX/10 && digit > MAX%10) {
            return 0
        }
        if res < MIN/10 || (res == MIN/10 && digit < MIN%10) {
            return 0
        }
        res = (res * 10) + digit
    }

    return res
}
