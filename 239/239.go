package main

/*
Вам дан массив целых чисел nums, есть скользящее окно размера k, которое движется из самого левого края массива в самый правый. 
В окне вы можете увидеть только k чисел. Каждый раз скользящее окно перемещается вправо на одну позицию.
Вернуть максимумы скользящих окон.

Пример 1:
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Объяснение: 
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

Пример 2:
Input: nums = [1], k = 1
Output: [1]

Ограничения:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length

*/

type deque struct {
	indexes []int
}

// func (d *deque) pushFirst(i int) {
// 	d.indexes = append([]int{i}, d.indexes...)
// }

func (d *deque) getFirst() int {
	return d.indexes[0]
}

func (d *deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *deque) pushLast(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *deque) empty() bool {
	return len(d.indexes) == 0
}


func maxSlidingWindow(nums []int, k int) []int {
    if k == 1 {
		return nums
	}
	res := make([]int, 0)
	dq := &deque{}
	l, r := 0, 0
	for ; r < len(nums); r++ {
		for !dq.empty() && nums[dq.getLast()] < nums[r] {
			dq.popLast()
		}
		dq.pushLast(r)

		if l > dq.getFirst() {
			dq.popFirst()
		}
		if r+1>=k {
			res = append(res, nums[dq.getFirst()])
			l++
		}
	}
	return res
}
