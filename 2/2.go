package main

/*
Вам даны два непустых связанных списка, представляющих два неотрицательных целых числа. 
Цифры хранятся в обратном порядке, и каждый из их узлов содержит одну цифру. Сложите два числа и верните сумму в виде связанного списка.
Вы можете предположить, что эти два числа не содержат ведущих нулей, кроме самого числа 0.

Пример 1:
https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Пояснение: 342+465=807.

Пример 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Пример 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:
Количество узлов в каждом связанном списке находится в диапазоне [1, 100].
0 <= Node.val <= 9
Гарантируется, что список представляет число, не имеющее ведущих нулей.

*/

type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
	tail := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		carry = val / 10
		val %= 10
		tail.Next = &ListNode{Val: val}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		tail = tail.Next
	}
	return dummy.Next
}
