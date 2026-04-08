package main

import "math"

/*
Реализуй pow(x, n) — функцию возведения x в степень n (то есть x^n).

Пример 1:
Вход: x = 2.00000, n = 10
Выход: 1024.00000

Пример 2:
Вход: x = 2.10000, n = 3
Выход: 9.26100

Пример 3:
Вход: x = 2.00000, n = -2
Выход: 0.25000
Пояснение: 2^(-2) = 1/2^2 = 1/4 = 0.25

Ограничения:
-100.0 < x < 100.0
-2^31 <= n <= 2^31-1
n — целое число.
Либо x ≠ 0, либо n > 0.
-10^4 <= x^n <= 10^4
*/

func myPow(x float64, n int) float64 {
    if x == 0 {
        return 0
    }
    if n == 0 {
        return 1
    }

    res := 1.0
    power := int(math.Abs(float64(n)))

    for power > 0 {
        if power&1 != 0 {
            res *= x
        }
        x *= x
        power >>= 1
    }

    if n >= 0 {
        return res
    }
    return 1 / res
}
