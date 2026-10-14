package main

import "sort"

/*
Вам дан список авиабилетов,
где tickets[i] = [from_i, to_i] обозначает аэропорты вылета и прибытия
одного рейса.
Восстановите маршрут в правильном порядке и верните его.

Все билеты принадлежат человеку, который вылетает из аэропорта «JFK»,
поэтому маршрут должен начинаться с «JFK».
Если существует несколько допустимых маршрутов,
верните маршрут с наименьшим лексическим порядком при чтении как одной строки.

Например, маршрут ["JFK", "LGA"] имеет меньший лексический порядок,
чем ["JFK", "LGB"].

Вы можете предположить, что все билеты составляют
хотя бы один действительный маршрут.
Вы должны использовать все билеты один и только один раз.

Пример 1:
Вход: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Выход: ["JFK","MUC","LHR","SFO","SJC"]

Пример 2:
Вход: tickets =
[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Выход: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Пояснение: Другой возможный маршрут — ["JFK","SFO","ATL","JFK","ATL","SFO"],
но он больше в лексическом порядке.

Ограничения:
1 <= tickets.length <= 300
tickets[i].length == 2
from_i.length == 3
to_i.length == 3
from_i и to_i состоят из заглавных английских букв.
from_i != to_i
*/

// findItinerary строит лексически минимальный эйлеров маршрут из JFK.
// Алгоритм Хиерхольцера использует билет как отдельное направленное ребро.
// E = len(tickets); при трёхбуквенных названиях время в среднем O(E log E).
// Рабочая память O(E): граф и стек DFS; результат хранит E + 1 строку.
// Строки ссылаются на названия из tickets, их байтовые буферы не копируются.
// tickets сортируется на месте; порядок элементов входного среза меняется.
func findItinerary(tickets [][]string) []string {
    // Граф: город → список направлений
    adj := make(map[string][]string)
    
    // Сравнение упорядочивает сначала вылет, затем прилёт за O(1) времени
    // и памяти: оба названия по контракту имеют длину три байта.
    sort.Slice(tickets, func(i, j int) bool {
        if tickets[i][0] != tickets[j][0] {
            return tickets[i][0] < tickets[j][0]
        }
        return tickets[i][1] < tickets[j][1]
    })
    
    // Обратный обход сортировки даёт у каждого src убывающий список dst.
    // Поэтому удаление с конца выбирает наименьшее доступное направление.
    // Повторные одинаковые билеты остаются отдельными элементами списка.
    for i := len(tickets) - 1; i >= 0; i-- {
        src := tickets[i][0]
        dst := tickets[i][1]
        adj[src] = append(adj[src], dst)
    }
    
    res := []string{}
    
    var dfs func(string)
    // dfs расходует исходящие рёбра и дописывает src в обратный маршрут.
    // За все вызовы время в среднем O(E), стек O(E), добавляется E + 1 строка.
    // Изменяет захваченные adj и res; один город может посещаться много раз.
    dfs = func(src string) {
        // Проверка непустого списка делает последний индекс допустимым.
        for len(adj[src]) > 0 {
            // Удаляем билет до рекурсии: даже при цикле он не повторится.
            dst := adj[src][len(adj[src])-1]
            adj[src] = adj[src][:len(adj[src])-1]
            dfs(dst)
        }
        // Постпорядок склеивает вложенные обходы, а не фиксирует жадный путь.
        // Например, тупиковый рейс из JFK окажется в конце готового маршрута.
        res = append(res, src)
    }
    
    // Начинаем с JFK
    dfs("JFK")
    
    // Реверсим результат (обратный обход → прямой порядок)
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
//             adj[src] = append(adj[src][:i],
//                 append([]string{v}, adj[src][i:]...)...)
//             res = res[:len(res)-1]
//         }
        
//         return false
//     }
    
//     dfs("JFK")
//     return res
// }