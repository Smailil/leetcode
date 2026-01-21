package main

/*
Дан массив целых чисел представляет дневную температуру, верните ответ массива, такой, что answer[i] — это количество дней, 
которое вам нужно подождать после i-го дня, чтобы получить более высокую температуру. 
Если нет будущего дня, для которого это возможно, вместо этого сохраните answer[i] == 0.

Пример 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Пример 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Пример 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]

Ограничения:
1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100

*/

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	answer := make([]int, len(temperatures))
	for i := range temperatures {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if temperatures[i] > temperatures[top] {
				answer[top] = i - top
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return answer
}