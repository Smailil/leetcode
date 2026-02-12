package main

/*
В этой задаче дерево представляет собой неориентированный граф, связный и не имеющий циклов.
Вам дан граф, который начинался как дерево с n узлами, помеченными от 1 до n, и добавленным одним дополнительным ребром. 
Добавленное ребро имеет две разные вершины, выбранные от 1 до n, и не является уже существующим ребром. 
Граф представлен как массив ребер длины n, где ребра [i] = [ai, bi] указывают на то, что между узлами ai и bi в графе есть ребро.
Возврайте ребро, которое можно удалить, чтобы полученный граф представлял собой дерево из n узлов. 
Если существует несколько ответов, верните ответ, который встречается последним во входных данных.

Пример 1:

Input: edges = [[1,2],[1,3],[2,3]]
Output: [2,3]

Пример 2:

Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]
 
Ограничения:
n == edges.length
3 <= n <= 1000
edges[i].length == 2
1 <= ai < bi <= edges.length
ai != bi
Повторяющихся ребер нет.
Данный граф связен.
*/

func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    adj := make([][]int, n+1)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    visit := make([]bool, n+1)
    cycle := make(map[int]struct{})
    cycleStart := -1

    var dfs func(node, par int) bool
    dfs = func(node, par int) bool {
        if visit[node] {
            cycleStart = node
            return true
        }

        visit[node] = true
        for _, nei := range adj[node] {
            if nei == par {
                continue
            }
            if dfs(nei, node) {
                if cycleStart != -1 {
                    cycle[node] = struct{}{}
                }
                if node == cycleStart {
                    cycleStart = -1
                }
                return true
            }
        }
        return false
    }

    dfs(1, -1)

    for i := len(edges) - 1; i >= 0; i-- {
        u, v := edges[i][0], edges[i][1]
		if _, ok := cycle[u]; ok {
			if _, ok := cycle[v]; ok {
				return []int{u, v}
			}
		}
    }
    return []int{}
}
