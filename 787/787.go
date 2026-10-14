package main

import "math"

/*
Имеется n городов, соединенных некоторым количеством рейсов.
Дан массив flights, где flights[i] = [from_i, to_i, price_i] означает,
что существует рейс из города from_i в город to_i стоимостью price_i.

Также даны три целых числа src, dst и k.
Верните минимальную стоимость пути из src в dst с не более чем k остановками.
Если такого маршрута нет, верните -1.

Пример 1:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png
Вход: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src
= 0, dst = 3, k = 1
Выход: 700
Объяснение:
Граф показан выше.
Оптимальный путь с не более чем 1 остановкой от города 0 до 3 отмечен красным и
стоит 100 + 600 = 700.
Обратите внимание, что путь через города [0,1,2,3] дешевле, но недействителен,
поскольку использует 2 остановки

Пример 2:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png
Вход: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Выход: 200
Объяснение:
Граф показан выше.
Оптимальный путь с не более чем 1 остановкой от города 0 до 2 отмечен красным и
стоит 100 + 100 = 200.

Пример 3:
https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png
Вход: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Выход: 500
Объяснение:
Граф показан выше.
Оптимальный путь без остановок из города 0 в город 2 отмечен красным и стоит
500.

Ограничения:
2 <= n <= 100
0 <= flights.length <= (n * (n - 1) / 2)
flights[i].length == 3
0 <= from_i, to_i < n
from_i != to_i
1 <= price_i <= 10^4
Нескольких рейсов между двумя городами не будет.
0 <= src, dst, k < n
src != dst
*/

// findCheapestPrice делает k+1 слоя релаксации Bellman–Ford: слой допускает
// не более заданного числа рёбер; k остановок означают k+1 перелётов.
// Время O(n+(k+1)*(n+E)), E=len(flights); рабочая память O(n) для двух слоёв.
// Результат int O(1); flights не меняется, индексы верны по условию.
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prices := make([]int, n)
    for i := range prices {
        prices[i] = math.MaxInt32
    }
    // Нулевой слой: src достижим без рёбер, остальные цены бесконечны.
    prices[src] = 0

    for range(k+1) {
        tmpPrices := make([]int, n)
        // Сохраняем пути с меньшим числом рёбер: это предел «не более».
        copy(tmpPrices, prices)

        // Переходы читают только прежний слой. Чтение нового позволило бы
        // пройти цепочку рёбер за одну итерацию и нарушить предел остановок.
        for _, flight := range flights {
            s, d, p := flight[0], flight[1], flight[2]
            // Не прибавляем к sentinel. Путь до k+1 рёбер стоит не больше
            // 10^6 по ограничениям, поэтому реальная цена безопасна для int32.
            if prices[s] == math.MaxInt32 {
                continue
            }
            if prices[s] + p < tmpPrices[d] {
                tmpPrices[d] = prices[s] + p
            }
        }
        // Следующая итерация допускает ещё один перелёт сверх этого слоя.
        prices = tmpPrices
    }

    if prices[dst] == math.MaxInt32 {
        return -1
    }
    return prices[dst]
}
