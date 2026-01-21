package main

/*
Дан заголовок связанного списка, определите, есть ли в связанном списке цикл.
В связанном списке существует цикл, если в списке есть некоторый узел, к которому можно снова добраться, непрерывно следуя по следующему указателю. 
Внутри pos используется для обозначения индекса узла, к которому подключен следующий указатель хвоста. Обратите внимание, что pos не передается в качестве параметра.
Возвращайте true, если в связанном списке есть цикл. В противном случае верните false.

Пример 1:
https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png
Input: head = [3,2,0,-4], pos = 1
Output: true
Пояснение: В связанном списке есть цикл, хвост которого соединяется с 1-м узлом (с индексом 0).

Пример 2:
https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png
Input: head = [1,2], pos = 0
Output: true
Пояснение: В связанном списке есть цикл, хвост которого соединяется с 0-м узлом.

Пример 3:
https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png
Input: head = [1], pos = -1
Output: false
Объяснение: В связанном списке нет цикла.

Ограничения:
Количество узлов в списке находится в диапазоне [0, 10^4].
-10^5 <= Node.val <= 10^5
pos имеет значение -1 или допустимый индекс в связанном списке.
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
