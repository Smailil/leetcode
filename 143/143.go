package main

/*
Вам дан заголовок односвязного списка. Список можно представить в виде:
L0 → L1 → … → Ln - 1 → Ln

Измените порядок списка, придав ему следующий вид:
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

Вы не можете изменять значения в узлах списка. Изменять можно только сами узлы.

Пример 1:
https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Пример 2:
https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

Ограничения:
Количество узлов в списке находится в диапазоне [1, 5 *10^4].
1 <= Node.val <= 1000

*/

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	var prev *ListNode
	curr := head2
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	head2 = prev
	for head2 != nil {
		nxt1, nxt2 := head.Next, head2.Next
		head.Next = head2
		head2.Next = nxt1
		head, head2 = nxt1, nxt2
	}
}
