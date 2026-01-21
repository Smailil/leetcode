package main

/*
Дан корень двоичного дерева поиска и целое число k, верните k-е наименьшее значение (с индексом 1) из всех значений узлов в дереве.

Пример 1:
https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg
Input: root = [3,1,4,null,2], k = 1
Output: 1

Пример 2:
https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
 
Ограничения:
Число узлов в дереве равно n.
1 <= k <= n <= 10^4
0 <= Node.val <= 10^4

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    stack := make([]*TreeNode, 0)
	n := 0
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n += 1 
		if n == k {
			return curr.Val
		}
		curr = curr.Right
	}
	return 0
}
