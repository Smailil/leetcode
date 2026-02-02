package main

/*
Вам дана сетка размером m x n, инициализированная этими тремя возможными значениями.

	-1 Стена или препятствие.
	0 Ворота.
	INF Бесконечность означает пустую комнату. Мы используем значение 2^31 - 1 = 2147483647. 
	для представления INF, поскольку вы можете предположить, что расстояние до ворот меньше 2147483647.
Заполните каждую пустую комнату расстоянием до ближайших ворот.
Если невозможно добраться до ворот, их следует заполнить INF.

Пример 1:

Input: rooms = [[2147483647,-1,0,2147483647],[2147483647,2147483647,2147483647,-1],
				[2147483647,-1,2147483647,-1],[0,-1,2147483647,2147483647]]
Output: [[3,-1,0,1],[2,2,1,-1],[1,-1,2,-1],[0,-1,3,4]]

Пример 2:
Input: rooms = [[-1]]
Output: [[-1]]

Ограничения:
m == rooms.length
n == rooms[i].length
1 <= m, n <= 250
rooms[i][j] имеет значение -1, 0 или 2^31 - 1.

*/

type cell struct {
	row int
	col int
}

func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
	visited := make(map[cell]struct{})
	queue := []cell{}

	addCell := func(r, c int) {
		if (r < 0 || r == rows || c < 0 || c == cols) {
			return
		}
		room := cell{r, c}
		if _, ok := visited[room]; ok {
			return
		}
		if grid[r][c] == -1 {
			return
		}
		visited[room] = struct{}{}
		queue = append(queue, room)
	}

	for r := range(rows) {
		for c := range(cols) {
			if grid[r][c] == 0 {
				treasure := cell{r, c}
				visited[treasure] = struct{}{}
				queue = append(queue, treasure)
			}
		}
	}
	distance := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		for range(levelSize) {
			current := queue[0]
			queue = queue[1:]
			grid[current.row][current.col] = distance
			addCell(current.row+1, current.col)
			addCell(current.row-1, current.col)
			addCell(current.row, current.col+1)
			addCell(current.row, current.col-1)
		}
		distance++
	}
}
