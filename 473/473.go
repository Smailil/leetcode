package main

import "sort"

/*
Дан целочисленный массив matchsticks, где matchsticks[i] — длина i-й спички.
Нужно использовать все спички, чтобы составить один квадрат.
Ломать спички нельзя, но можно соединять их встык,
и каждая спичка должна быть использована ровно один раз.

Верните true, если такой квадрат можно составить, и false в противном случае.

Пример 1:
Вход: matchsticks = [1,1,2,2,2]
Выход: true
Пояснение: Можно составить квадрат со стороной 2,
одна сторона квадрата получилась из двух спичек длиной 1.

Пример 2:
Вход: matchsticks = [3,3,3,3,4]
Выход: false
Пояснение: Нельзя найти способ составить квадрат из всех спичек.

Ограничения:
1 <= matchsticks.length <= 15
1 <= matchsticks[i] <= 10^8
*/

func makesquare(matchsticks []int) bool {
    sum := 0
    for _, m := range matchsticks {
        sum += m
    }
    if sum%4 != 0 {
        return false
    }
    
    target := sum / 4
    // Сортируем спички по убыванию для быстрого отсечения
    sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
    
    // Если самая длинная спичка больше стороны квадрата, ответ сразу false
    if matchsticks[0] > target {
        return false
    }

    sides := make([]int, 4)

    var dfs func(i int) bool
    dfs = func(i int) bool {
        // Базовый случай: все спички разложены. 
        // Благодаря проверкам ниже мы знаем, что ни одна сторона не превышает target,
        // а значит они все автоматически равны target.
        if i == len(matchsticks) {
            return true 
        }

        for j := range 4 {
            // Отсечение 1: не выходим за пределы целевой длины стороны
            if sides[j] + matchsticks[i] > target {
                continue
            }
            
            sides[j] += matchsticks[i]
            if dfs(i + 1) {
                return true
            }
            sides[j] -= matchsticks[i]
            
            // Отсечение 2: если сторона была пуста, и мы не смогли собрать квадрат,
            // то нет смысла пробовать другие пустые стороны (они эквивалентны).
            if sides[j] == 0 {
                break
            }
        }

        return false
    }

    return dfs(0)
}