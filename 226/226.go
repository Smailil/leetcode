package main

/*
Учитывая корень двоичного дерева, инвертируйте дерево и верните его корень.

Пример 1:
https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Пример 2:
https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg
Input: root = [2,1,3]
Output: [2,3,1]

Пример 3:
Input: root = []
Output: []

Ограничения:
Количество узлов в дереве находится в диапазоне [0, 100].
-100 <= Node.val <= 100

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
    	root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
