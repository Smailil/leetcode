package main

/*
Вам дан массив из k списков связанных списков, каждый связанный список отсортирован в порядке возрастания.
Объедините все связанные списки в один отсортированный связанный список и верните его.

Пример 1:
Объяснение: Связанные списки:
[
 1->4->5,
 1->3->4,
 2->6
]
объединение их в один отсортированный связанный список:
1->1->2->3->4->4->5->6

Пример 2:
Input: lists = []
Output: []

Пример 3:
Input: lists = [[]]
Output: []

Ограничения:
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] сортируется по возрастанию.
Сумма lists[i].length не будет превышать 10^4.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		mergedLists := make([]*ListNode, 0, (len(lists)+1)/2)
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}
	return lists[0]
}
