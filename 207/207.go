package main

/*
Всего вам нужно пройти numCourses курсов, помеченных от 0 до numCourses-1.
Вам предоставляется массив предварительных требований,
где prerequisites[i] = [a_i, b_i] указывает,
что вы должны сначала пройти курс b_i, если хотите пройти курс a_i.
Например, пара [0, 1] указывает, что для прохождения курса 0
необходимо сначала пройти курс 1.
Верните true, если вы можете закончить все курсы.
В противном случае верните false.

Пример 1:
Вход: numCourses = 2, prerequisites = [[1,0]]
Выход: true
Объяснение: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вам необходимо закончить курс 0.
Так что это возможно.

Пример 2:
Вход: numCourses = 2, prerequisites = [[1,0],[0,1]]
Выход: false
Объяснение: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вы должны закончить курс 0,
а чтобы пройти курс 0, вы также должны закончить курс 1.
Так что это невозможно.

Ограничения:
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= a_i, b_i < numCourses
Все пары prerequisites[i] уникальны.
*/

// canFinish проверяет отсутствие цикла в графе "курс -> его предпосылки".
// V=numCourses, E=len(prerequisites): среднее время O(V+E) с int-key map.
// Рабочая память O(V+E) для графа, O(V) для visiting/стека; ответ bool O(1).
// Входные пары не меняются: очищается только локальный preMap.
func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
    for i := 0; i < numCourses; i++ {
        preMap[i] = []int{}
    }
    // Пара [crs,pre] имеет два элемента, индексы курсов допустимы по условию.
    // Направление ведёт к требованиям курса, а не к его зависимым курсам.
    for _, prereq := range prerequisites {
        crs, pre := prereq[0], prereq[1]
        preMap[crs] = append(preMap[crs], pre)
    }

	visiting := make(map[int]struct{})
	var dfs func(int) bool
	// dfs проверяет требования crs; visiting содержит текущий путь рекурсии.
	// Суммарно все вызовы: среднее время O(V+E), стек/visiting O(V).
	// Граф preMap занимает O(V+E); повторно доказанный курс проверяем за O(1).
	// Меняет visiting и очищает успешные списки preMap, не входные пары.
	dfs = func(crs int) bool {
		// Возврат к узлу активного пути означает цикл, включая самозависимость.
		if _, ok := visiting[crs]; ok {
			return false
		}
		// Либо требований изначально нет, либо они уже успешно проверены.
		// Оба случая позволяют завершить crs без нового обхода.
		if len(preMap[crs]) == 0 {
			return true
		}
		visiting[crs] = struct{}{}
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		// Все требования доказаны: убираем узел из активного пути и запоминаем
		// успех пустым списком. Его рёбра больше не будут сканироваться.
		delete(visiting, crs)
		preMap[crs] = []int{}
		return true
	}
	// Проверяем все компоненты, в том числе не связанные с первым курсом.
	// При цикле dfs может оставить visiting непустым, но сразу завершаем вызов.
	for crs := range numCourses {
		if !dfs(crs) {
			return false
		}
	}
	return true
}
