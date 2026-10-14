package main

/*
У вас есть граф из n узлов, помеченных от 0 до n - 1.
Вам дано целое число n и список рёбер, где edges[i] = [a_i, b_i]
означает наличие ненаправленного ребра между узлами a_i и b_i.
Верните true, если рёбра данного графа составляют допустимое дерево,
и false в противном случае.

Пример 1:
https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0200-0299/0261.Graph%20Valid%20Tree/images/tree1-graph.jpg
Вход: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Выход: true

Пример 2:
https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0200-0299/0261.Graph%20Valid%20Tree/images/tree2-graph.jpg
Вход: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Выход: false

Ограничения:
1 <= n <= 2000
0 <= edges.length <= 5000
edges[i].length == 2
0 <= a_i, b_i < n
a_i != b_i
Никаких петель и повторяющихся рёбер нет.
*/

// validTree требует одновременно отсутствия цикла и связности графа.
// E=len(edges): среднее время O(n+E), рабочая память O(n+E),
// включая adjacency, map и стек; ответ bool O(1). edges не меняется.
// По условию n>=1, вершина 0 существует, все пары/индексы корректны.
func validTree(n int, edges [][]int) bool {
	adj := make([][]int, n)
	// Неориентированное ребро хранится в обоих списках, всего 2*E записей.
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visit := make(map[int]struct{})
	var dfs func(int, int) bool
	// dfs помечает компоненту и отвергает повторный вход не через родителя.
	// За весь обход: среднее время O(n+E), visit/стек O(n).
	// Списки adj уже учтены снаружи; меняет visit, не edges/adj.
	dfs = func(node, prev int) bool {
		// В неориентированном графе после пропуска родительского ребра
		// возвращение в любую посещённую вершину свидетельствует о цикле.
		if _, ok := visit[node]; ok {
			return false
		}
		visit[node] = struct{}{}
		for _, nei := range adj[node] {
			// Обратная запись того же ребра не образует цикл из двух вершин.
			// Отсутствие повторных рёбер делает пропуск по prev корректным.
			if nei == prev {
				continue
			}
			if !dfs(nei, node) {
				return false
			}
		}
		return true
	}

	// -1 не является вершиной: у стартового узла нет родительского ребра.
	// DFS одной компоненты недостаточно: требуем достижимость всех n узлов.
	// Для n=1 без рёбер оба условия истинны — это одноузловое дерево.
	return dfs(0, -1) && len(visit) == n
}
