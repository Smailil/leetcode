package main

/*\
У вас есть граф из n узлов, помеченных от 0 до n - 1. 
Вам дано целое число n и список ребер, где ребра [i] = [ai, bi] указывают на то, что между узлами ai и bi в графе есть ненаправленное ребро.
Возвращает true, если ребра данного графа составляют допустимое дерево, и false в противном случае.
 

Example 1:
https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0200-0299/0261.Graph%20Valid%20Tree/images/tree1-graph.jpg
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true

Example 2:
https://fastly.jsdelivr.net/gh/doocs/leetcode@main/solution/0200-0299/0261.Graph%20Valid%20Tree/images/tree2-graph.jpg
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false

Constraints:
1 <= n <= 2000
0 <= edges.length <= 5000
edges[i].length == 2
0 <= ai, bi < n
ai != bi
Никаких петель и повторяющихся ребер нет.

*/

func validTree(n int, edges [][]int) bool {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visit := make(map[int]struct{})
	var dfs func(int, int) bool
	dfs = func(node, prev int) bool {
		if _, ok := visit[node]; ok {
			return false
		}
		visit[node] = struct{}{}
		for _, nei := range adj[node] {
			if nei == prev {
				continue
			}
			if !dfs(nei, node) {
				return false
			}
		}
		return true
	}

	return dfs(0, -1) && len(visit) == n
}
