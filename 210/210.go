package main

/*
Всего вам нужно пройти numCourses курсов, помеченных от 0 до numCourses -1. 
Вам предоставляется массив предварительных требований, где preреквизиты[i] = [ai, bi] указывают, что вы должны сначала пройти курс bi, если хотите пройти курс ai.
Например, пара [0, 1] указывает, что для прохождения курса 0 необходимо сначала пройти курс 1.
Верните порядок курсов, которые вам необходимо пройти, чтобы закончить все курсы. Если правильных ответов много, верните любой из них.
Если пройти все курсы невозможно, верните пустой массив.

Пример 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Объяснение: Всего нужно пройти 2 курса. Чтобы пройти курс 1, вы должны закончить курс 0. Поэтому правильный порядок курсов — [0,1].

Пример 2:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Пояснение: Всего необходимо пройти 4 курса. Чтобы пройти курс 3, вам необходимо пройти оба курса 1 и 2. Оба курса 1 и 2 следует пройти после окончания курса 0.
Таким образом, один правильный порядок курса — [0,1,2,3]. Другой правильный порядок — [0,2,1,3].

Пример 3:
Input: numCourses = 1, prerequisites = []
Output: [0]

Ограничения:
1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
Все пары [ai, bi] различны.

*/

func findOrder(numCourses int, prerequisites [][]int) []int {
    prereq := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        prereq[i] = []int{}
	}
	for _, pair := range prerequisites {
		crs, pre := pair[0], pair[1]
        prereq[crs] = append(prereq[crs], pre)
	}

	output := []int{}
    visit := make(map[int]struct{})
    cycle := make(map[int]struct{})

    var dfs func(int) bool
    dfs = func(crs int) bool {
        if _, ok := cycle[crs]; ok {
            return false
        }
        if _, ok := visit[crs]; ok {
            return true
        }

        cycle[crs] = struct{}{}
        for _, pre := range prereq[crs] {
            if !dfs(pre) {
                return false
            }
        }
        delete(cycle, crs)
        visit[crs] = struct{}{}
        output = append(output, crs)
        return true
    }

    for i := range numCourses {
        if !dfs(i) {
            return []int{}
        }
    }

    return output
}
