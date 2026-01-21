package main

/*
Вам даны заголовки двух отсортированных связанных списков list1 и list2.
Объедините два списка в один отсортированный связанный список и верните заголовок нового отсортированного связанного списка.
Новый список должен состоять из узлов из списка list1 и list2.

Пример 1:
Input: list1 = [1,2,4], list2 = [1,3,5]
Output: [1,1,2,3,4,5]

Пример 2:
Input: list1 = [], list2 = [1,2]
Output: [1,2]

Пример 3:
Input: list1 = [], list2 = []
Output: []

Ограничения:
0 <= Длина каждого списка <= 100.
-100 <= Node.val <= 100

*/

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next

	}
	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}
	return dummy.Next
}
