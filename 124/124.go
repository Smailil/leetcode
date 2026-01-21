package main

/*
Путь в бинарном дереве — это последовательность узлов, в которой каждая пара соседних узлов в последовательности имеет соединяющее их ребро.
Узел может появляться в последовательности не более одного раза. Обратите внимание, что путь не обязательно должен проходить через корень.
Сумма путей пути — это сумма значений узлов в пути.
Учитывая корень двоичного дерева, верните максимальную сумму путей любого непустого пути.

Пример 1:
https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg
Input: root = [1,2,3]
Output: 6
Пояснение: Оптимальный путь — 2 -> 1 -> 3 с суммой путей 2 + 1 + 3 = 6.

Пример 2:
https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Пояснение: Оптимальный путь — 15 -> 20 -> 7 с суммой путей 15 + 20 + 7 = 42.

Ограничения:
Количество узлов в дереве находится в диапазоне [1, 3 *10^4].
-1000 <= Node.val <= 1000

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := root.Val
	var dfs func(root *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(dfs(node.Left), 0)
		rightMax := max(dfs(node.Right), 0)
		res = max(res, node.Val+leftMax+rightMax)
		return node.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return res
}
