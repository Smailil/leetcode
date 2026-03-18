package main

import "sort"

/*
Вам дан список авиабилетов, где tickets[i] = [fromi, toi] обозначают аэропорты вылета и прибытия одного рейса.
Восстановите маршрут в правильном порядке и верните его.

Все билеты принадлежат человеку, который вылетает из аэропорта «JFK», поэтому маршрут должен начинаться с «JFK».
Если существует несколько допустимых маршрутов, вам следует вернуть маршрут с наименьшим лексическим порядком при чтении как одной строки.

Например, маршрут ["JFK", "LGA"] имеет меньший лексический порядок, чем ["JFK", "LGB"].

Вы можете предположить, что все билеты составляют хотя бы один действительный маршрут. Вы должны использовать все билеты один и только один раз.

Example 1:

Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Output: ["JFK","MUC","LHR","SFO","SJC"]

Example 2:

Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"] but it is larger in lexical order.

Ограничения:
1 <= tickets.length <= 300
tickets[i].length == 2
fromi.length == 3
toi.length == 3
fromi и toi состоят из заглавных английских букв.
fromi != toi

*/

func findItinerary(tickets [][]string) []string {
    // Граф: город → список направлений
    adj := make(map[string][]string)
    
    // Сортируем билеты лексикографически
    sort.Slice(tickets, func(i, j int) bool {
        if tickets[i][0] != tickets[j][0] {
            return tickets[i][0] < tickets[j][0]
        }
        return tickets[i][1] < tickets[j][1]
    })
    
    // Строим граф в обратном порядке (для pop() с конца)
    for i := len(tickets) - 1; i >= 0; i-- {
        src := tickets[i][0]
        dst := tickets[i][1]
        adj[src] = append(adj[src], dst)
    }
    
    res := []string{}
    
    // DFS функция (Hierholzer's algorithm)
    var dfs func(string)
    dfs = func(src string) {
        // Пока есть исходящие рейсы
        for len(adj[src]) > 0 {
            // Берём последний пункт назначения (pop)
            dst := adj[src][len(adj[src])-1]
            adj[src] = adj[src][:len(adj[src])-1]
            dfs(dst)
        }
        // Добавляем город после обработки всех рёбер (post-order)
        res = append(res, src)
    }
    
    // Начинаем с JFK
    dfs("JFK")
    
    // Реверсим результат (post-order → прямой порядок)
    for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }
    
    return res
}

// func findItinerary(tickets [][]string) []string {
//     // Строим граф
//     adj := make(map[string][]string)
//     for _, ticket := range tickets {
//         adj[ticket[0]] = append(adj[ticket[0]], ticket[1])
//     }
    
//     // Сортируем направления для каждого города
//     for src := range adj {
//         sort.Strings(adj[src])
//     }
    
//     res := []string{"JFK"}
    
//     var dfs func(string) bool
//     dfs = func(src string) bool {
//         if len(res) == len(tickets)+1 {
//             return true
//         }
        
//         if _, exists := adj[src]; !exists {
//             return false
//         }
        
//         // Копируем список (чтобы безопасно модифицировать)
//         temp := make([]string, len(adj[src]))
//         copy(temp, adj[src])
        
//         for i, v := range temp {
//             // Удаляем рейс
//             adj[src] = append(adj[src][:i], adj[src][i+1:]...)
//             res = append(res, v)
            
//             if dfs(v) {
//                 return true
//             }
            
//             // Бэктрекинг: восстанавливаем
//             adj[src] = append(adj[src][:i], append([]string{v}, adj[src][i:]...)...)
//             res = res[:len(res)-1]
//         }
        
//         return false
//     }
    
//     dfs("JFK")
//     return res
// }