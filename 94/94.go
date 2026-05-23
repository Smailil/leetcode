package main

/*
Дан корень бинарного дерева — верните значения его узлов в порядке
симметричного (inorder) обхода.

Пример 1:
Вход: root = [1,null,2,3]
Выход: [1,3,2]
Пояснение:
https://assets.leetcode.com/uploads/2024/08/29/screenshot-2024-08-29-202743.png

Пример 2:
Вход: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Выход: [4,2,6,5,7,1,3,9,8]
Пояснение:
https://assets.leetcode.com/uploads/2024/08/29/tree_2.png

Пример 3:
Вход: root = []
Выход: []

Пример 4:
Вход: root = [1]
Выход: [1]

Ограничения:
Количество узлов дерева находится в диапазоне [0, 100].
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

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        res = append(res, node.Val)
        inorder(node.Right)
    }
    inorder(root)
    return res
}