package main

/*
Дан массив целых чисел, представляющих высоту столбца гистограммы, где ширина каждого столбца равна 1, 
верните площадь самого большого прямоугольника в гистограмме.

Пример 1:
Input: heights = [2,1,5,6,2,3]
Output: 10
https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg
Пояснение: Выше представлена гистограмма, ширина каждого столбца которой равна 1.
Самый большой прямоугольник показан в красной области, его площадь = 10 единиц.

Пример 2:
Input: heights = [2,4]
Output: 4
https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg

Ограничения:
1 <= heights.length <= 105
0 <= heights[i] <= 104

*/

func largestRectangleArea(heights []int) int {
    stack := make([][2]int, 0)
	maxAria := 0
	for i := range heights {
		newI := i
		for len(stack) > 0 && stack[len(stack)-1][1] > heights[i] {
			indx, height := stack[len(stack)-1][0], stack[len(stack)-1][1]
			maxAria = max(maxAria, (i - indx)*height)
			stack = stack[:len(stack)-1]
			newI = indx
		}
		stack = append(stack, [2]int{newI, heights[i]})
	}
	for _, s := range stack {
		indx, height := s[0], s[1]
		maxAria = max(maxAria, (len(heights) - indx)*height)
	}
	return maxAria
}

func main() {
	h := []int{2,1,5,6,2,3}
	largestRectangleArea(h)
}