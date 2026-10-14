package main

/*
Дан массив целых чисел, представляющих высоту столбца гистограммы,
где ширина каждого столбца равна 1,
верните площадь самого большого прямоугольника в гистограмме.

Пример 1:
Вход: heights = [2,1,5,6,2,3]
Выход: 10
https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg
Пояснение: Выше представлена гистограмма,
ширина каждого столбца которой равна 1.
Самый большой прямоугольник показан в красной области, его площадь = 10 единиц.

Пример 2:
Вход: heights = [2,4]
Выход: 4
https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg

Ограничения:
1 <= heights.length <= 10^5
0 <= heights[i] <= 10^4
*/

// largestRectangleArea завершает прямоугольники монотонным стеком.
// n = len(heights); время O(n) амортизированно: запись входит/выходит раз.
// Рабочая память O(n), результат — число; вход не изменяется.
// Площадь <= 10^5*10^4 = 10^9, умножение помещается в 32-битный int.
func largestRectangleArea(heights []int) int {
    stack := make([][2]int, 0)
	maxAria := 0
	// [start,height] означает начало допустимого интервала этой высоты.
	// Высоты стека не убывают; равные записи разрешены кодом.
	for i := range heights {
		newI := i
		// Меньшая колонка i завершает все более высокие кандидаты слева.
		for len(stack) > 0 && stack[len(stack)-1][1] > heights[i] {
			indx, height := stack[len(stack)-1][0], stack[len(stack)-1][1]
			// Все колонки [indx,i) не ниже height, ширина равна i-indx;
			// сама i уже ниже, поэтому не входит в прямоугольник.
			maxAria = max(maxAria, (i - indx)*height)
			stack = stack[:len(stack)-1]
			// Нынешняя меньшая высота может занять весь вынутый интервал.
			newI = indx
		}
		stack = append(stack, [2]int{newI, heights[i]})
	}
	// Оставшиеся кандидаты не встретили меньшей высоты справа:
	// их правая исключённая граница — len(heights), ширина n-indx.
	for _, s := range stack {
		indx, height := s[0], s[1]
		maxAria = max(maxAria, (len(heights) - indx)*height)
	}
	return maxAria
}

// main запускает пример; возвращённая площадь здесь не печатается.
// Для фиксированных 6 столбцов время O(1) и рабочая память O(1).
// Для примера из n столбцов вызов largestRectangleArea занимает
// O(n) времени и O(n) рабочей памяти стека; создание h тоже O(n).
func main() {
	h := []int{2,1,5,6,2,3}
	largestRectangleArea(h)
}