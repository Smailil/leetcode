package main

/*
Дан корень двоичного дерева, верните его максимальную глубину.
Максимальная глубина двоичного дерева — это количество узлов на самом длинном пути от корневого узла до самого дальнего листового узла.

Пример 1:

Input: root = [3,9,20,null,null,15,7]
Output: 3

Пример 2:
Input: root = [1,null,2]
Output: 2

Ограничения:
Количество узлов в дереве находится в диапазоне [0, 10^4].
-100 <= Node.val <= 100

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
