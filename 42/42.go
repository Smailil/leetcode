package main

/*
Даны n неотрицательных целых чисел, представляющих карту высот, где ширина каждой полосы равна 1, 
вычислите, сколько воды она может удержать после дождя.

Пример 1:
https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Пояснение: Приведенная выше карта высот (черная секция) представлена ​​массивом [0,1,0,2,1,0,1,3,2,1,2,1]. 
В данном случае задерживается 6 единиц дождевой воды (синяя секция).

Пример 2:
Input: height = [4,2,0,3,2,5]
Output: 9

Ограничения:
n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

*/

func trap(height []int) int {
	res := 0
	if len(height) != 0 {
    	l, r := 0, len(height) - 1
		maxL, maxR := height[l], height[r]
		for l < r {
			if maxL < maxR {
				l++
				res += max(0, maxL-height[l])
				maxL = max(maxL, height[l])
			} else {
				r--
				res += max(0, maxR-height[r])
				maxR = max(maxR, height[r])
			}
		}
	}
	return res
}
