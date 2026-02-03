package main

/*
Вам дана сетка размером m x n, где каждая ячейка может иметь одно из трех значений:
0 представляет пустую ячейку,
1 представляет собой свежий апельсин или
2 представляет собой гнилой апельсин.

Каждую минуту любой свежий апельсин, примыкающий к гнилому апельсину в четырех направлениях, становится гнилым.
Возвращайте минимальное количество минут, которое должно пройти до тех пор, пока ни в одной ячейке не появится свежий апельсин. Если это невозможно, верните -1.

Пример 1:
https://assets.leetcode.com/uploads/2019/02/16/oranges.png
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Пример 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Пояснение: Апельсин в левом нижнем углу (строка 2, столбец 0) никогда не бывает гнилым, поскольку гниение происходит только в 4 направлениях.

Пример 3:
Input: grid = [[0,2]]
Output: 0
Пояснение: Поскольку на минуте 0 свежих апельсинов уже нет, ответ будет всего лишь 0.

Ограничения:
m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] равна 0, 1 или 2.

*/

type orangesCoords struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	q := []orangesCoords{}
	time, fresh := 0, 0
	rows, cols := len(grid), len(grid[0])

	for r := range rows {
		for c := range cols {
			switch grid[r][c] {
			case 1:
				fresh++
			case 2:
				q = append(q, orangesCoords{r, c})
			}
		}
	}
	for len(q) != 0 && fresh > 0 {
		currentRotten := len(q)
		for range currentRotten {
			cur := q[0]
			q = q[1:]
			for _, d := range directions {
				r, c := d[0]+cur.row, d[1]+cur.col
				if r < 0 || r == rows ||
					c < 0 || c == cols ||
					grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				q = append(q, orangesCoords{r, c})
				fresh--
			}
		}
		time++
	}
	if fresh == 0 {
		return time
	} else {
		return -1
	}
}
