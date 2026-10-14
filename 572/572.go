package main

/*
Даны корни двух двоичных деревьев root и subRoot.
Верните true, если в root существует поддерево
с той же структурой и значениями узлов, что и subRoot,
и false в противном случае.
Поддерево состоит из узла дерева и всех потомков этого узла.
Само дерево также можно рассматривать как собственное поддерево.

Пример 1:
https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg
Вход: root = [3,4,5,1,2], subRoot = [4,1,2]
Выход: true

Пример 2:
https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg
Вход: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Выход: false

Ограничения:
Количество узлов в дереве root находится в диапазоне [1, 2000].
Количество узлов в дереве subRoot находится в диапазоне [1, 1000].
-10^4 <= root.val <= 10^4
-10^4 <= subRoot.val <= 10^4
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree проверяет совпадение структуры и значений двух деревьев.
// p — число сравнённых непустых пар, d — глубина проверенной части.
// Время O(p+1), рабочая память стека O(d+1), результат bool занимает O(1).
// Узлы и ссылки входных деревьев не изменяются.
func isSameTree(tree1, tree2 *TreeNode) bool {
	// Две пустые ветви совпадают; одна пустая ветвь даст false ниже.
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 != nil && tree2 != nil && tree1.Val == tree2.Val {
		// После неудачи слева правая пара уже не проверяется из-за &&.
		return isSameTree(tree1.Left, tree2.Left) &&
			isSameTree(tree1.Right, tree2.Right)
	}
	return false
}

// isSubtree ищет узел root, чьё полное поддерево совпадает с subRoot.
// Для непустых деревьев с n и m узлами худшее время O(n*m):
// каждый из n кандидатов может потребовать сравнения до m узлов.
// Рабочая память O(h), h — высота root: сравнение идёт ниже кандидата,
// поэтому вместе с путём поиска стек не выходит за высоту root.
// Результат bool занимает O(1); входные деревья не изменяются.
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// Пустой образец считается поддеревом любого дерева.
	// Если один корень пуст, ответ получается за O(1) времени и памяти.
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	// Сначала проверяем текущее полное поддерево, затем дочерние узлы.
	if isSameTree(root, subRoot) {
		return true
	} else {
		// Успех слева прекращает поиск: || не вызывает правую ветвь.
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}
}
