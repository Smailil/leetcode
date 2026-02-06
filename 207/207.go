package main

/*
Всего вам нужно пройти numCourses курсов, помеченных от 0 до numCourses-1. 
Вам предоставляется массив предварительных требований, где prerequisites[i] = [ai, bi] указывают, что вы должны сначала пройти курс bi, если хотите пройти курс ai.
Например, пара [0, 1] указывает, что для прохождения курса 0 необходимо сначала пройти курс 1.
Верните true, если вы можете закончить все курсы. В противном случае верните false.

Пример 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Объяснение: Всего нужно пройти 2 курса. 
Чтобы пройти курс 1, вам необходимо закончить курс 0. Так что это возможно.

Пример 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Объяснение: Всего нужно пройти 2 курса. 
Чтобы пройти курс 1, вы должны закончить курс 0, а чтобы пройти курс 0, вы также должны закончить курс 1. Так что это невозможно.

Constraints:
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
Все пары prerequisites[i] уникальны.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }
    for _, prereq := range prerequisites {
        crs, pre := prereq[0], prereq[1]
        preMap[crs] = append(preMap[crs], pre)
    }

	visiting := make(map[int]struct{})
	var dfs func(int) bool
	dfs = func(crs int) bool {
		if _, ok := visiting[crs]; ok {
			return false
		}
		if len(preMap[crs]) == 0 {
			return true
		}
		visiting[crs] = struct{}{}
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visiting, crs)
		preMap[crs] = []int{}
		return true
	}
	for crs := range numCourses {
		if !dfs(crs) {
			return false
		}
	}
	return true
}
