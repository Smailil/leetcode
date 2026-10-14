package main

/*
Даны два целочисленных массива preorder и inorder,
где preorder — это прямой обход двоичного дерева,
а inorder — это симметричный обход того же дерева.
Постройте и верните двоичное дерево.

Пример 1:
https://assets.leetcode.com/uploads/2021/02/19/tree.jpg
Вход: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Выход: [3,9,20,null,null,15,7]

Пример 2:
Вход: preorder = [-1], inorder = [-1]
Выход: [-1]

Ограничения:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder и inorder состоят из уникальных значений.
Каждое значение inorder также появляется в preorder.
preorder гарантированно является прямым обходом дерева.
inorder гарантированно является симметричным обходом дерева.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree восстанавливает дерево из согласованных уникальных обходов.
// Корень берётся первым в preorder, его позиция в inorder делит поддеревья.
// Среднее время O(n), n=len(preorder), при O(1) поиске целых ключей в map.
// Рабочая память O(n+h): карта индексов и стек высоты h; результат O(n) узлов.
// Входные массивы не изменяются; узлы результата создаются заново.
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Строим map: значение -> индекс в inorder.
	// Уникальность из условия делает разделение поддеревьев однозначным.
	inorderIndexMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}

	// Передаём включённые границы, не создавая срезы поддиапазонов.
	// Обычный срез тоже не копировал бы элементы исходного массива.
	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	// build восстанавливает поддерево из диапазонов одинаковой длины.
	// Среднее время O(k), k — число его узлов; рабочая память стека O(h).
	// Общая карта не копируется; результат O(k) новых узлов.
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		// Пустой диапазон означает отсутствующего ребёнка и защищает индекс.
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[preStart]
		rootInorderIndex := inorderIndexMap[rootVal]
		// Слева от корня в inorder ровно rootInorderIndex-inStart узлов.
		// Столько же элементов после корня preorder образуют левое поддерево.
		leftSubtreeSize := rootInorderIndex - inStart
		// Корень не входит в дочерние диапазоны; остаток справа — правое дерево.
		return &TreeNode{
			Val:   rootVal,
			Left:  build(preStart+1, preStart+leftSubtreeSize, inStart, rootInorderIndex-1),
			Right: build(preStart+leftSubtreeSize+1, preEnd, rootInorderIndex+1, inEnd),
		}
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
