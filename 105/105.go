package main

/*
Даны два целочисленных массива preorder и inorder, где preorder — это предварительный обход двоичного дерева,
а inorder — это неупорядоченный обход того же дерева, постройте и верните двоичное дерево.

Пример 1:
https://assets.leetcode.com/uploads/2021/02/19/tree.jpg
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Пример 2:
Input: preorder = [-1], inorder = [-1]
Output: [-1]


Ограничения:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder и inorder состоят из уникальных значений.
Каждое значение inorder также появляется в предварительном порядке.
preorder гарантированно является предварительным обходом дерева.
inorder гарантированно будет неупорядоченным обходом дерева.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// Строим map: значение -> индекс в inorder
	inorderIndexMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}

	// Используем вспомогательную функцию с индексами, чтобы избежать копирования слайсов
	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[preStart]
		rootInorderIndex := inorderIndexMap[rootVal]
		leftSubtreeSize := rootInorderIndex - inStart
		return &TreeNode{
			Val:   rootVal,
			Left:  build(preStart+1, preStart+leftSubtreeSize, inStart, rootInorderIndex-1),
			Right: build(preStart+leftSubtreeSize+1, preEnd, rootInorderIndex+1, inEnd),
		}
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
