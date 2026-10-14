package main

/*
Дан корень root двоичного дерева.
Верните значения его узлов в постфиксном порядке (postorder).

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
// postorderTraversal собирает значения в порядке левое, правое, корень.
// Время O(n), n — число узлов; рабочая память O(h) для рекурсии высоты h.
// Память результата O(n); дерево не изменяется, сохраняются копии значений.
// В этом файле TreeNode объявлен лишь в комментарии: нужно внешнее определение.
func postorderTraversal(root *TreeNode) []int {
    res := []int{}
    var postorder func(node *TreeNode)
    // postorder дописывает поддерево в общий res за O(k), k — число его узлов.
    // Рабочая память O(h) для стека; добавленная часть результата O(k).
    postorder = func(node *TreeNode) {
        // Пустое поддерево завершает эту ветку без добавления значений.
        if node == nil {
            return
        }
        postorder(node.Left)
        postorder(node.Right)
        // Корень добавляется после обоих поддеревьев.
        res = append(res, node.Val)
    }
    postorder(root)
    return res
}