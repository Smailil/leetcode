package main

/*
Дано двоичное дерево, определите, сбалансировано ли оно по высоте.

Пример 1:
https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg
Input: root = [3,9,20,null,null,15,7]
Output: true

Пример 2:
https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false

Пример 3:
Input: root = []
Output: true

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

func isBalanced(root *TreeNode) bool {
	var checkBalance func(*TreeNode) balanceInfo
	checkBalance = func(node *TreeNode) balanceInfo {
		if node == nil {
			return balanceInfo{true, 0}
		}

		leftInfo := checkBalance(node.Left)
		rightInfo := checkBalance(node.Right)

		heightDiff := leftInfo.height - rightInfo.height
		heightDiff = max(heightDiff, -heightDiff)

		isBalanced := leftInfo.isBalanced &&
			rightInfo.isBalanced &&
			(heightDiff <= 1)

		return balanceInfo{
			isBalanced, 1 + max(leftInfo.height, rightInfo.height),
		}
	}

	return checkBalance(root).isBalanced
}
