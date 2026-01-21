package main

/*
Дано двоичное дерево поиска (BST), найдите узел наименьшего общего предка (LCA) из двух заданных узлов в BST.
Согласно определению LCA в Википедии: «Наименьший общий предок определяется между двумя узлами p и q как самый нижний узел в T,
который имеет как p, так и q в качестве потомков (где мы позволяем узлу быть потомком самого себя)».

Пример 1:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Пояснение: LCA узлов 2 и 8 равен 6.

Пример 2:

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Объяснение: LCA узлов 2 и 4 равен 2, поскольку согласно определению LCA узел может быть потомком самого себя.

Пример 3:
Input: root = [2,1], p = 2, q = 1
Output: 2

Ограничения:
Количество узлов в дереве находится в диапазоне [2, 10^5].
-10^9 <= Node.val <= 10^9
Все Node.val уникальны.
p != q
p и q будут существовать в BST.

*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}
	return cur
}
