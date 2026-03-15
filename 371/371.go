package main

/*
Даны два целых числа a и b, верните сумму двух целых чисел без использования операторов + и -.

Пример 1:
Input: a = 1, b = 2
Output: 3

Пример 2:
Input: a = 2, b = 3
Output: 5

Ограничения:
-1000 <= a, b <= 1000

*/

func getSum(a int, b int) int {
    mask := 0xFFFFFFFF   // Маска 32 бит: эмулирует переполнение unsigned int
    maxInt := 0x7FFFFFFF // Макс. положительное 32-битное знаковое число

    for b != 0 {
        carry := (a & b) << 1  // Биты переноса
        a = (a ^ b) & mask     // Сумма без переноса, обрезка до 32 бит
        b = carry & mask       // Перенос, обрезка до 32 бит
    }

    if a <= maxInt {
        return a  // Положительное: бит 31 = 0
    }
    return ^(a ^ mask)  // Отрицательное: конвертация из 32-бит two's complement в Go int
}