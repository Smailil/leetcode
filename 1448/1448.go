package main

/*
Дан корень двоичного дерева, узел X в дереве называется хорошим, если на пути от корня до X нет узлов со значением, большим, чем X.
Возвращает количество хороших узлов в бинарном дереве.

Пример 1:
https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png
Input: root = [3,1,4,3,null,1,5]
Output: 4
Пояснение: Узлы, выделенные синим цветом, являются исправными.
Корневой узел (3) всегда является хорошим узлом.
Узел 4 -> (3,4) — максимальное значение на пути, начиная с корня.
Узел 5 -> (3,4,5) — максимальное значение на пути.
Узел 3 -> (3,1,3) — максимальное значение на пути.

Пример 2:
https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png
Input: root = [3,3,null,4,2]
Output: 3
Пояснение: Узел 2 -> (3, 3, 2) не подходит, потому что «3» выше него.

Пример 3:
Input: root = [1]
Output: 1
Объяснение: Root считается хорошим.

Ограничения:
Количество узлов в бинарном дереве находится в диапазоне [1, 10^5].
Значение каждого узла находится в диапазоне [-10^4, 10^4].

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, prevMax int) int
	dfs = func(node *TreeNode, prevMax int) int {
		if node == nil {
			return 0
		}
		res := 0
		if prevMax <= node.Val {
			res = 1
			prevMax = node.Val
		}
		return res + dfs(node.Left, prevMax) + dfs(node.Right, prevMax)
	}
	return dfs(root, root.Val)
}
