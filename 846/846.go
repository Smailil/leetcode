package main

/*
У Алисы есть карты, которые она хочет распределить по группам размера groupSize.
В каждой группе значения карт должны быть последовательными.
Даны целочисленный массив hand, где hand[i] — значение i-й карты,
и целое число groupSize.
Верните true, если карты можно распределить таким образом, иначе false.

Пример 1:
Вход: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Выход: true
Пояснение: Карты Алисы можно распределить как [1,2,3],[2,3,4],[6,7,8]

Пример 2:
Вход: hand = [1,2,3,4,5], groupSize = 4
Выход: false
Пояснение: Карты Алисы нельзя разложить на группы по 4 штуки.

Ограничения:
1 <= hand.length <= 10^4
0 <= hand[i] <= 10^9
1 <= groupSize <= hand.length
*/

import (
    "sort"
)

// isNStraightHand начинает группы с минимальной ещё оставшейся карты.
// Все её freq копий обязаны начинать группы: меньших свободных карт нет.
// Среднее время O(n+d*log(d+1)), n=len(hand), d — число разных карт.
// Рабочая память O(d), результат bool O(1); hand не меняется, сортируются keys.
// Успешный шаг внутреннего цикла расходует хотя бы одну карту: таких шагов
// не больше n, даже если разные группы частично перекрываются.
func isNStraightHand(hand []int, groupSize int) bool {
    // groupSize положителен по условию; число карт должно делиться на него.
    if len(hand)%groupSize != 0 {
        return false
    }

    count := make(map[int]int)
    for _, n := range hand {
        count[n]++
    }

    keys := make([]int, 0, len(count))
    for k := range count {
        keys = append(keys, k)
    }
    // Порядок обхода map не определён, поэтому отдельно сортируем ключи.
    sort.Ints(keys)

    for _, key := range keys {
        freq := count[key]
        // Эта карта могла полностью уйти в группы с меньшим началом.
        if freq == 0 {
            continue
        }

        // Каждое из freq обязательных начал требует по карте каждого значения.
        // key+groupSize не больше 10^9+10^4 и безопасен для int32.
        for i := key; i < key+groupSize; i++ {
            // Отсутствующий ключ даёт0; заменить недостающую карту нельзя.
            if count[i] < freq {
                return false
            }
            count[i] -= freq
        }
    }

    return true
}

// import (
//     "container/heap"
// )

// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
//     *h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[0 : n-1]
//     return x
// }

// func isNStraightHand(hand []int, groupSize int) bool {
//     if len(hand)%groupSize != 0 {
//         return false
//     }

//     count := make(map[int]int)
//     for _, n := range hand {
//         count[n]++
//     }

//     minH := make(IntHeap, 0, len(count))
//     for k := range count {
//         minH = append(minH, k)
//     }
//     heap.Init(&minH)

//     for minH.Len() > 0 {
//         first := minH[0]
//         freq := count[first] // Сколько групп начинается с этого элемента

//         for i := first; i < first+groupSize; i++ {
//             if count[i] < freq {
//                 return false
//             }
//             count[i] -= freq
//             if count[i] == 0 {
//                 if i != minH[0] {
//                     return false
//                 }
//                 heap.Pop(&minH)
//             }
//         }
//     }

//     return true
// }
