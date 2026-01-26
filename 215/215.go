package main

import "container/heap"

/*
Дан целочисленный массив nums и целое число k, верните k-й по величине элемент массива.
Обратите внимание, что это k-й по величине элемент в отсортированном порядке, а не k-й отдельный элемент.
Можно ли решить задачу без сортировки?

Пример 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Пример 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4

Ограничения:
1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

*/

type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0:n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    mh := make(MinHeap, k)
    copy(mh, nums[:k])
    heap.Init(&mh)
    
    for i := k; i < len(nums); i++ {
        if nums[i] > mh[0] {
            heap.Pop(&mh)
            heap.Push(&mh, nums[i])
        }
    }
    return mh[0]
}