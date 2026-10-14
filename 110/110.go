package main

/*
Дано двоичное дерево, определите, сбалансировано ли оно по высоте.

Пример 1:
https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg
Вход: root = [3,9,20,null,null,15,7]
Выход: true

Пример 2:
https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg
Вход: root = [1,2,2,3,3,null,null,4,4]
Выход: false

Пример 3:
Вход: root = []
Выход: true

Ограничения:
Количество узлов в дереве находится в диапазоне [0, 5000].
-10^4 <= Node.val <= 10^4
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type balanceInfo struct {
	isBalanced bool
	height     int
}

// isBalanced проверяет баланс дерева одним обходом снизу вверх.
// n — число узлов, h — высота: время O(n), рабочая память O(h) для стека.
// Результат — один bool, O(1); исходное дерево не меняется.
func isBalanced(root *TreeNode) bool {
	var checkBalance func(*TreeNode) balanceInfo
	// checkBalance возвращает высоту поддерева и баланс всех его узлов.
	// Для s узлов и высоты h: время O(s), дополнительный стек O(h),
	// результат balanceInfo занимает O(1). Дерево/внешнее состояние не меняет.
	checkBalance = func(node *TreeNode) balanceInfo {
		// Пустое поддерево сбалансировано и имеет высоту 0.
		if node == nil {
			return balanceInfo{true, 0}
		}

		// Оба поддерева обходятся до проверки &&: раннего выхода из DFS нет.
		leftInfo := checkBalance(node.Left)
		rightInfo := checkBalance(node.Right)

		// max(d,-d) даёт |d|. Высоты <= 5000, поэтому отрицание разности
		// и увеличение высоты безопасны для int.
		heightDiff := leftInfo.height - rightInfo.height
		heightDiff = max(heightDiff, -heightDiff)

		// Нужен баланс обоих детей и разность их высот не выше 1.
		isBalanced := leftInfo.isBalanced &&
			rightInfo.isBalanced &&
			(heightDiff <= 1)

		// Высота — один текущий узел плюс более высокая ветвь.
		return balanceInfo{
			isBalanced, 1 + max(leftInfo.height, rightInfo.height),
		}
	}

	return checkBalance(root).isBalanced
}
