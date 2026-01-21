package main

/*
Дан корень двоичного дерева, верните длину диаметра дерева.
Диаметр двоичного дерева — это длина самого длинного пути между любыми двумя узлами дерева. Этот путь может проходить через корень, а может и не проходить.
Длина пути между двумя узлами представлена ​​количеством ребер между ними.

Пример 1:
Input: root = [1,2,3,4,5]
Output: 3
Пояснение: 3 — это длина пути [4,2,1,3] или [5,2,1,3].

Пример 2:
Input: root = [1,2]
Output: 1

Ограничения:
Количество узлов в дереве находится в диапазоне [1, 10^4].
-100 <= Node.val <= 100

*/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    var dfs func(*TreeNode) int
    dfs = func(curr *TreeNode) int {
        if curr == nil {
            return 0
        }
        left := dfs(curr.Left)
        right := dfs(curr.Right)
		res = max(res, left+right)
        return max(left, right) + 1
    }
    dfs(root)
    return res
}