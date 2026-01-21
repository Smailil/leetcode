package main

/*
Учитывая заголовок связанного списка, удалите n-й узел из конца списка и верните его заголовок.

Пример 1:

Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Пример 2:
Input: head = [1], n = 1
Output: []

Пример 3:
Input: head = [1,2], n = 1
Output: [1]

Ограничения:
Количество узлов в списке равно sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	l, r := dummy, head
	for ; n > 0; n-- {
		r = r.Next
	}
	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
