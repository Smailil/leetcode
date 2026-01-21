package main

/*
Даны корни двух двоичных деревьев p и q, напишите функцию, проверяющую, совпадают они или нет.
Два двоичных дерева считаются одинаковыми, если они структурно идентичны и узлы имеют одинаковое значение.

Пример 1:
https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg
Input: p = [1,2,3], q = [1,2,3]
Output: true

Пример 2:
https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg
Input: p = [1,2], q = [1,null,2]
Output: false

Пример 3:
https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg
Input: p = [1,2,1], q = [1,1,2]
Output: false

Ограничения:
Количество узлов в обоих деревьях находится в диапазоне [0, 100].
-10^4 <= Node.val <= 10^4

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
