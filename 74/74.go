package main

/*
Вам дана целочисленная матрица размера m x n со следующими двумя свойствами:
Каждая строка сортируется в порядке неубывания.
Первое целое число каждой строки больше последнего целого числа предыдущей строки.
Дана целочисленная цель, верните true, если цель находится в матрице, или false в противном случае.

Вы должны написать решение с временной сложностью O(log(m *n)).

Пример 1:
https://assets.leetcode.com/uploads/2020/10/05/mat.jpg
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Пример 2:
https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false

Ограничения:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10^4 <= matrix[i][j], target <= 10^4

*/

func searchMatrix(matrix [][]int, target int) bool {
    l, r := 0, len(matrix)*len(matrix[0]) - 1
	for l <= r {
		mid := (l+r)/2
		m, n := mid/len(matrix[0]), mid%len(matrix[0])
		if target > matrix[m][n] {
			l = mid + 1
		} else if target < matrix[m][n] {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
