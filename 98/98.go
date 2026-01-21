package main

import "math"

/*
Дан корень двоичного дерева, определите, является ли оно допустимым двоичным деревом поиска (BST).
Действительный BST определяется следующим образом:
Левое поддерево узла содержит только узлы с ключами, строго меньшими, чем ключ узла.
Правое поддерево узла содержит только узлы с ключами, строго превышающими ключ узла.
И левое, и правое поддеревья также должны быть бинарными деревьями поиска.

Пример 1:
https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg
Input: root = [2,1,3]
Output: true

Пример 2:
https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg
Input: root = [5,1,4,null,null,3,6]
Output: false
Объяснение: Значение корневого узла равно 5, а значение его правого дочернего узла равно 4.

Ограничения:
Количество узлов в дереве находится в диапазоне [1, 10^4].
-2^31 <= Node.val <= 2^31 - 1

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
    var valid func(node *TreeNode, left int, right int) bool
	valid = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}
		if !(left < node.Val && node.Val < right) {
			return false
		}
		return valid(node.Left, left, node.Val) && valid(node.Right, node.Val, right)
	}
	return valid(root, math.MinInt, math.MaxInt)
}
