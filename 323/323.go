package main

/*
У вас есть граф из n узлов. Вам задано целое число n и массив ребер, где ребра [i] = [aᵢ, bᵢ] указывают на то, что между aᵢ и bᵢ в графе есть ребро.
Возвращайте количество связных компонентов в графе.

Example 1:
Input:
n = 5, edges = [[0,1],[1,2],[3,4]]
Output: 2

Example 2:
Input:
n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
Output: 1

Constraints:
1 <= n <= 2000
1 <= edges.length <= 5000
edges[i].length == 2
0 <= aᵢ <= bᵢ < n
aᵢ != bᵢ
Повторяющихся ребер нет.

*/

func countComponents(n int, edges [][]int) int {
    adj := make([][]int, n)
	visit := make([]bool, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	var dfs func(int)
	dfs = func(node int) {
		for _, nei := range adj[node] {
			if !visit[nei] {
				visit[nei] = true
				dfs(nei)
			}
		}
	}

	res := 0
	for node := 0; node < n; node++ {
		if !visit[node] {
			visit[node] = true
			dfs(node)
			res++
		}
	}
	return res
}

