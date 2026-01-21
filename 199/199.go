package main

/*
Дан корень двоичного дерева, представьте, что вы стоите на его правой стороне, и верните значения узлов, которые вы видите, упорядоченные сверху вниз.

Пример 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
Explanation:
https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png

Пример 2:
Input: root = [1,2,3,4,null,null,null,5]
Output: [1,3,4,5]
Explanation:
https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png

Пример 3:
Input: root = [1,null,3]
Output: [1,3]

Пример 4:
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

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		rightNode := &TreeNode{}
		qLen := len(q)
		for i := 0; i<qLen ; i++ {
			rightNode = q[0]
			q = q[1:]
			if rightNode.Left != nil {
				q = append(q, rightNode.Left)
			}
			if rightNode.Right != nil {
				q = append(q, rightNode.Right)
			}
		}
		res = append(res, rightNode.Val)
	}
	return res
}
