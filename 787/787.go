package main

import "math"

/*
Имеется n городов, соединенных некоторым количеством рейсов.
Вам дан массив рейсов, где flights[i] = [fromi, toi,pricei] указывают, что существует рейс из города fromi в город toi по pricei.

Вам также даны три целых числа src, dst и k, которые возвращают самую дешевую цену от src до dst с не более чем k остановками. Если такого маршрута нет, верните -1.

Пример 1:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png
Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
Output: 700
Объяснение:
График показан выше.
Оптимальный путь с не более чем 1 остановкой от города 0 до 3 отмечен красным и стоит 100 + 600 = 700.
Обратите внимание, что путь через города [0,1,2,3] дешевле, но недействителен, поскольку использует 2 остановки

Пример 2:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
Объяснение:
График показан выше.
Оптимальный путь с не более чем 1 остановкой от города 0 до 2 отмечен красным и стоит 100 + 100 = 200.

Пример 3:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
Объяснение:
График показан выше.
Оптимальный путь без остановок из города 0 в город 2 отмечен красным и стоит 500.

Ограничения:
2 <= n <= 100
0 <= flights.length <= (n * (n - 1) / 2)
flights[i].length == 3
0 <= fromi, toi < n
fromi != toi
1 <= pricei <= 10^4
Нескольких рейсов между двумя городами не будет.
0 <= src, dst, k < n
src != dst

*/

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prices := make([]int, n)
    for i := range prices {
        prices[i] = math.MaxInt32
    }
    prices[src] = 0

    for range(k+1) {
        tmpPrices := make([]int, n)
        copy(tmpPrices, prices)

        for _, flight := range flights {
            s, d, p := flight[0], flight[1], flight[2]
            if prices[s] == math.MaxInt32 {
                continue
            }
            if prices[s] + p < tmpPrices[d] {
                tmpPrices[d] = prices[s] + p
            }
        }
        prices = tmpPrices
    }

    if prices[dst] == math.MaxInt32 {
        return -1
    }
    return prices[dst]
}
