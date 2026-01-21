package main

/*
Дано начало заголовка односвязного списка, переверните список и верните начало нового списка.
Пример 1:
Input: head = [0,1,2,3]
Output: [3,2,1,0]

Пример 2:
Input: head = []
Output: []

Ограничения:
0 <= Длина списка <= 1000.
-1000 <= Node.val <= 1000

*/

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

