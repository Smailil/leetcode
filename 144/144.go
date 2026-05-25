package main

/*
Дано корневое дерево, верните прямой обход (preorder) значений его узлов.

Пример 1:
Вход: root = [1,null,2,3]
Выход: [1,2,3]
Пояснение:
https://assets.leetcode.com/uploads/2024/08/29/screenshot-2024-08-29-202743.png

Пример 2:
Вход: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Выход: [1,2,4,5,6,7,3,8,9]
Пояснение:
https://assets.leetcode.com/uploads/2024/08/29/tree_2.png

Пример 3:
Вход: root = []
Выход: []

Пример 4:
Вход: root = [1]
Выход: [1]

Ограничения:
Количество узлов в дереве — в диапазоне [0, 100].
-100 <= Node.val <= 100
*/

/**
 * Определение узла бинарного дерева.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res := []int{}
    var preorder func(node *TreeNode)
    preorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        res = append(res, node.Val)
        preorder(node.Left)
        preorder(node.Right)
    }
    preorder(root)
    return res
}
