package main

import "sort"

/*
Дан список учетных записей: каждый элемент accounts[i] - это список строк,
где первый элемент accounts[i][0] - имя, а остальные элементы - email-адреса,
относящиеся к этой учетной записи.

Нужно объединить эти учетные записи. Две учетные записи точно принадлежат
одному человеку, если у них есть общий email. Обрати внимание: даже если
у двух учетных записей одинаковое имя, они могут принадлежать разным людям,
потому что у разных людей могут совпадать имена. Изначально у человека может
быть любое количество учетных записей, но все его учетные записи точно имеют
одно и то же имя.

После объединения верни учетные записи в следующем формате: первый элемент
каждой учетной записи - имя, а остальные элементы - email-адреса
в отсортированном порядке. Сами учетные записи можно вернуть в любом порядке.

Пример 1:
Вход:
accounts = [
	["John","johnsmith@mail.com","john_newyork@mail.com"],
	["John","johnsmith@mail.com","john00@mail.com"],
	["Mary","mary@mail.com"],
	["John","johnnybravo@mail.com"],
]
Выход:
[
	["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],
	["Mary","mary@mail.com"],
	["John","johnnybravo@mail.com"],
]
Пояснение:
Первая и вторая учетные записи John принадлежат одному человеку,
потому что у них есть общий email "johnsmith@mail.com".
Третья учетная запись John и учетная запись Mary принадлежат разным людям,
потому что ни один из их email-адресов не используется в других учетных записях.
Мы можем вернуть эти списки в любом порядке, например ответ
[['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']]
тоже будет принят.

Пример 2:
Вход:
accounts = [
	["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],
	["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],
	["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],
	["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],
	["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"],
]
Выход:
[
	["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],
	["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],
	["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],
	["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],
	["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"],
]

Ограничения:
1 <= accounts.length <= 1000
2 <= accounts[i].length <= 10
1 <= accounts[i][j].length <= 30
accounts[i][0] состоит из английских букв.
accounts[i][j] (для j > 0) является корректным email.
*/

// Подход: построение неориентированного графа и поиск связных компонент (Connected Components).
// Вершины графа — это уникальные email-адреса.
// Ребра соединяют email-адреса, принадлежащие одному аккаунту.
// Каждая связная компонента графа представляет собой объединенный аккаунт одного пользователя.
func accountsMerge(accounts [][]string) [][]string {
    // emailIdx: маппинг email -> уникальный целочисленный ID (индекс вершины в графе).
    emailIdx := make(map[string]int)
    // emails: обратный маппинг ID -> email (нужен, чтобы по ID вершины получить сам email).
    emails := []string{}
    // emailToAcc: маппинг ID email-адреса -> ID исходного аккаунта (индекс в accounts).
    // Нужен, чтобы в конце достать имя пользователя для итоговой группы.
    emailToAcc := make(map[int]int)

    m := 0 // Счетчик уникальных email-адресов (он же — количество вершин в графе)
    
    // Шаг 1: Назначаем уникальные ID всем email-адресам.
    for accId, account := range accounts {
        for i := 1; i < len(account); i++ { // начинаем с 1, так как 0 — это имя
            email := account[i]
            if _, exists := emailIdx[email]; !exists {
                emails = append(emails, email)     // сохраняем email
                emailIdx[email] = m                // присваиваем ID
                emailToAcc[m] = accId              // запоминаем, из какого аккаунта этот email
                m++                                // увеличиваем счетчик вершин
            }
        }
    }

    // Шаг 2: Строим граф (список смежности).
    adj := make([][]int, m)
    for i := range adj {
        adj[i] = []int{} // инициализируем пустые срезы для каждой вершины
    }
    
    // Соединяем email-адреса внутри каждого аккаунта ребрами.
    for _, account := range accounts {
        // Оптимизация: чтобы не делать O(N^2) ребер для аккаунта с N email-ами,
        // мы соединяем каждый email только с предыдущим (цепочкой: 1-2, 2-3, 3-4).
        // Этого достаточно, чтобы все email-адреса аккаунта оказались в одной связной компоненте.
        for i := 2; i < len(account); i++ {
            id1 := emailIdx[account[i]]     // ID текущего email
            id2 := emailIdx[account[i-1]]   // ID предыдущего email
            
            // Добавляем ребра в обе стороны (граф неориентированный)
            adj[id1] = append(adj[id1], id2)
            adj[id2] = append(adj[id2], id1)
        }
    }

    // Шаг 3: Поиск связных компонент с помощью DFS (поиска в глубину).
    emailGroup := make(map[int][]string) // ID аккаунта (корня компоненты) -> список email-адресов
    visited := make([]bool, m)           // массив посещенных вершин графа

    // Рекурсивная функция DFS (замыкание, имеет доступ к adj, visited, emails и emailGroup)
    var dfs func(node, accId int)
    dfs = func(node, accId int) {
        visited[node] = true // отмечаем вершину как посещенную
        // добавляем email текущей вершины в группу соответствующего аккаунта
        emailGroup[accId] = append(emailGroup[accId], emails[node])
        
        // рекурсивно обходим всех соседей
        for _, nei := range adj[node] {
            if !visited[nei] {
                dfs(nei, accId)
            }
        }
    }

    // Запускаем DFS для каждой непосещенной вершины
    for i := 0; i < m; i++ {
        if !visited[i] {
            // В качестве идентификатора группы (accId) передаем ID аккаунта,
            // к которому изначально принадлежал этот email (берем из emailToAcc).
            dfs(i, emailToAcc[i])
        }
    }

    // Шаг 4: Формируем итоговый результат.
    res := [][]string{}
    for accId, group := range emailGroup {
        name := accounts[accId][0]   // берем имя пользователя из исходного аккаунта
        sort.Strings(group)          // сортируем email-адреса (требование задачи)
        
        // создаем итоговый срез: [имя, email1, email2, ...]
        merged := append([]string{name}, group...)
        res = append(res, merged)
    }

    return res
}