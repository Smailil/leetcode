package main

/*
У Алисы есть некоторое количество карточек, и она хочет перегруппировать карточки в группы так, чтобы каждая группа имела размер groupSize и
состояла из последовательных карточек groмupSize.
Дана hand - целочисленный массив, где hand[i] — это значение, записанное на i-й карте, и целочисленный groupSize, 
верните true, если она может перегруппировать карты, или false в противном случае.
Example 1:
Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Пояснение: Руку Алисы можно переставить как [1,2,3],[2,3,4],[6,7,8]

Example 2:
Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Пояснение: Руку Алисы нельзя разложить на группы по 4 штуки.

Ограничения:
1 <= hand.length <= 10^4
0 <= hand[i] <= 10^9
1 <= groupSize <= hand.length

*/

import (
    "sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
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
    sort.Ints(keys)

    for _, key := range keys {
        freq := count[key]
        if freq == 0 {
            continue
        }

        for i := key; i < key+groupSize; i++ {
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
