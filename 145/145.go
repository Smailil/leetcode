package main

/*
Дано корневое значение бинарного дерева;
вернуть обход его узлов в постфиксном порядке.

Пример 1:
Вход: root = [1,null,2,3]
Выход: [3,2,1]
Пояснение:
https://assets.leetcode.com/uploads/2024/08/29/screenshot-2024-08-29-202743.png

Пример 2:
Вход: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Выход: [4,6,7,5,2,9,8,3,1]
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
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    var postorder func(node *TreeNode)
    postorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        postorder(node.Left)
        postorder(node.Right)
        res = append(res, node.Val)
    }
    postorder(root)
    return res
}