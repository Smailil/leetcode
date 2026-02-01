package main

/*
Дана ссылка на узел связного неориентированного графа.
Верните глубокую копию (клон) графика.
Каждый узел графа содержит значение (int) и список (List[Node]) своих соседей.
class Node {
    public int val;
    public List<Node> neighbors;
}
 

Формат тестового примера:
Для простоты значение каждого узла совпадает с индексом узла (с индексом 1). 
Например, первый узел с val == 1, второй узел с val == 2 и так далее. 
Граф представлен в тестовом примере с использованием списка смежности.
Список смежности — это набор неупорядоченных списков, используемый для представления конечного графа. 
Каждый список описывает набор соседей узла в графе.
Данный узел всегда будет первым узлом с val = 1.
Вы должны вернуть копию данного узла как ссылку на клонированный граф.

Пример 1:
https://assets.leetcode.com/uploads/2019/11/04/133_clone_graph_question.png
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Пояснение: В графе 4 узла.
Соседями 1-го узла (val = 1) являются 2-й узел (val = 2) и 4-й узел (val = 4).
Соседями 2-го узла (val = 2) являются 1-й узел (val = 1) и 3-й узел (val = 3).
Соседями 3-го узла (val = 3) являются 2-й узел (val = 2) и 4-й узел (val = 4).
Соседями четвертого узла (val = 4) являются 1-й узел (val = 1) и 3-й узел (val = 3).

Пример 2:
https://assets.leetcode.com/uploads/2020/01/07/graph.png
Input: adjList = [[]]
Output: [[]]
Объяснение: Обратите внимание, что входные данные содержат один пустой список. Граф состоит только из одного узла с val = 1 и не имеет соседей.

Пример 3:
Input: adjList = []
Output: []
Пояснение: Это пустой граф, в нем нет узлов.

Ограничения:
Количество узлов в графе находится в диапазоне [0, 100].
1 <= Node.val <= 100
Node.val уникален для каждого узла.
В графе нет повторяющихся ребер и петель.
Граф связен, и все узлы можно посетить, начиная с данного узла.

*/

type Node struct {
    Val int
    Neighbors []*Node
} 

// func cloneGraph(node *Node) *Node {
//     if node == nil {
// 		return nil
// 	}
//     visited := make([]*Node, 101)
//     var dfs func(*Node)
//     dfs = func(n *Node) {
//         newN := &Node{Val: n.Val}
//         visited[n.Val] = newN
//         for _, neig := range n.Neighbors {
//             if visited[neig.Val] == nil {
//                 dfs(neig)
//             }
//             newN.Neighbors = append(newN.Neighbors, visited[neig.Val])
//         }
//     }
//     dfs(node)
//     return visited[node.Val]
// }

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    visited := make(map[*Node]*Node) // map старый -> новый узел
    var dfs func(*Node) *Node
    
    dfs = func(n *Node) *Node {
        if cloned, ok := visited[n]; ok {
            return cloned
        }
        
        newN := &Node{Val: n.Val}
        visited[n] = newN
        
        for _, neighbor := range n.Neighbors {
            newN.Neighbors = append(newN.Neighbors, dfs(neighbor))
        }
        
        return newN
    }
    
    return dfs(node)
}
