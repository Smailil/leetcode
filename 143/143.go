package main

/*
Вам дана голова односвязного списка.
Список можно представить в виде:
L_0 → L_1 → … → L_{n-1} → L_n

Измените порядок списка, придав ему следующий вид:
L_0 → L_n → L_1 → L_{n-1} → L_2 → L_{n-2} → …

Вы не можете изменять значения в узлах списка.
Изменять можно только сами узлы.

Пример 1:
https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg
Вход: head = [1,2,3,4]
Выход: [1,4,2,3]

Пример 2:
https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg
Вход: head = [1,2,3,4,5]
Выход: [1,5,2,4,3]

Ограничения:
Количество узлов в списке находится в диапазоне [1, 5 *10^4].
1 <= Node.val <= 1000
*/

type ListNode struct {
    Val int
    Next *ListNode
}

// reorderList чередует узлы начала списка с узлами его конца.
// Находит середину, разрывает список, разворачивает вторую часть и сплетает.
// Время O(n), n — число узлов; рабочая память O(1), новых узлов нет.
// Связи Next исходного списка меняются; значения Val сохраняются.
// Требует непустой head, что гарантировано условием задачи.
func reorderList(head *ListNode)  {
    // fast впереди на один узел: первая половина не короче второй.
    slow, fast := head, head.Next
	// Обе проверки защищают fast.Next.Next от разыменования nil.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Разрыв исключает цикл при сплетении половин.
	head2 := slow.Next
	slow.Next = nil
	var prev *ListNode
	curr := head2
	// prev — развёрнутый префикс второй половины, curr — остаток.
	for curr != nil {
		// Сохраняем хвост до перезаписи ссылки Next.
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	head2 = prev
	// Первой половины хватает; при нечётной длине центр останется последним.
	for head2 != nil {
		// Сохраняем обе исходные связи перед сплетением пары узлов.
		nxt1, nxt2 := head.Next, head2.Next
		head.Next = head2
		head2.Next = nxt1
		head, head2 = nxt1, nxt2
	}
}
