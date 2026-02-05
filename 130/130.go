package main

/*
Вам дана матричная доска размером m x n, содержащая буквы «X» и «O», захватите области, которые окружены:

Соединение: ячейка соединяется с соседними ячейками по горизонтали или по вертикали.
Регион: Чтобы сформировать регион, соедините каждую ячейку «O».
Окружение: регион окружен ячейками «X», если вы можете соединить регион с ячейками «X», и ни одна из ячеек региона не находится на краю доски.
Чтобы захватить окруженный регион, замените все буквы «О» на «Х» на исходной доске. Вам не нужно ничего возвращать.

Пример 1:
Input: board = [["X","X","X","X"],
				["X","O","O","X"],
				["X","X","O","X"],
				["X","O","X","X"]]
Output: [["X","X","X","X"],
		 ["X","X","X","X"],
		 ["X","X","X","X"],
		 ["X","O","X","X"]]
Объяснение:
https://assets.leetcode.com/uploads/2021/02/19/xogrid.jpg
На приведенной выше диаграмме нижняя область не захвачена, поскольку она находится на краю доски и не может быть окружена.

Пример 2:
Input: board = [["X"]]
Output: [["X"]]

Ограничения:
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] имеет значение «X» или «O».

*/

func solve(board [][]byte)  {
    rows, cols := len(board), len(board[0])
	var lockCell func(int, int)
	lockCell = func(r, c int) {
		if r < 0 || c < 0 || r == rows ||
           c == cols || board[r][c] != 'O' {
            return
        }
        board[r][c] = 'L'
        lockCell(r+1, c)
        lockCell(r-1, c)
        lockCell(r, c+1)
        lockCell(r, c-1)
	}
	for r := range rows {
        if board[r][0] == 'O' {
            lockCell(r, 0)
        }
        if board[r][cols-1] == 'O' {
            lockCell(r, cols-1)
        }
    }

    for c := range cols {
        if board[0][c] == 'O' {
            lockCell(0, c)
        }
        if board[rows-1][c] == 'O' {
            lockCell(rows-1, c)
        }
    }

	for r := range rows {
		for c := range cols {
			switch board[r][c] {
			case 'O':
				board[r][c] = 'X'
			case 'L':
				board[r][c] = 'O'
			}
		}
	}
}
