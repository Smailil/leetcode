package main

/*
Дан корень двоичного дерева, верните порядок прохождения уровней значений его узлов. (т. е. слева направо, уровень за уровнем).

Пример 1:
https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Пример 2:
Input: root = [1]
Output: [[1]]

Пример 3:
Input: root = []
Output: []

Ограничения:
Количество узлов в дереве находится в диапазоне [0, 2000].
-1000 <= Node.val <= 1000

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		qLen := len(q)
		level := make([]int, 0)
		for i := 0; i < qLen; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}
