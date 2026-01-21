package main

/*
Даны корни двух двоичных деревьев root и subRoot, верните true,
если существует корневое поддерево с той же структурой и значениями узлов subRoot, и false в противном случае.
Поддерево двоичного дерева — это дерево, состоящее из узла дерева и всех потомков этого узла. Дерево-дерево также можно рассматривать как поддерево самого себя.

Example 1:
https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Example 2:
https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false

Ограничения:
Количество узлов в корневом дереве находится в диапазоне [1, 2000].
Количество узлов в подкорневом дереве находится в диапазоне [1, 1000].
-10^4 <= root.val <= 10^4
-10^4 <= subRoot.val <= 10^4

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 != nil && tree2 != nil && tree1.Val == tree2.Val {
		return isSameTree(tree1.Left, tree2.Left) &&
			isSameTree(tree1.Right, tree2.Right)
	}
	return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}
