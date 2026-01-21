package main

import "math"

/*
Квадратная тройка (a,b,c) — это тройка, где a, b и c — целые числа и a^2 + b^2 = c^2.
Дано целое число n, верните количество троек квадратов таких, что 1 <= a, b, c <= n.

Пример 1:
Input: n = 5
Output: 2
Пояснение: Квадратные тройки — это (3,4,5) и (4,3,5).

Example 2:
Input: n = 10
Output: 4
Пояснение: Квадратные тройки — это (3,4,5), (4,3,5), (6,8,10) и (8,6,10).
 
Ограничения:
1 <= n <= 250

*/

func countTriples(n int) int {
    count := 0
    nSquared := n * n
    
    for a := 1; a <= n; a++ {
        aSquared := a * a
        for b := a; b <= n; b++ {
            bSquared := b * b
            cSquared := aSquared + bSquared
            
            if cSquared > nSquared {
                break
            }
            
            c := int(math.Sqrt(float64(cSquared)))
            if c*c == cSquared {
                if a == b {
                    count++
                } else {
                    count += 2
                }
            }
        }
    }
    
    return count
}
