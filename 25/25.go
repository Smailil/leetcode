package main

/*
Дан заголовок связанного списка, поменяйте местами узлы списка k за раз и верните измененный список.
k — целое положительное число, меньшее или равное длине связанного списка. 
Если количество узлов не кратно k, то пропущенные узлы в конечном итоге должны остаться такими, какие они есть.
Вы не можете изменять значения в узлах списка, можно изменять только сами узлы.

Пример 1:
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]

Пример 2:
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]

Ограничения:
Количество узлов в списке равно n.
1 <= k <= n <= 5000
0 <= Node.val <= 1000

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{0, head}
	groupPrev := dummy

	for {
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break
		}
		groupNext := kth.Next
		prev, curr := groupNext, groupPrev.Next
		for curr != groupNext {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		next := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = next
	}
	return dummy.Next
}

