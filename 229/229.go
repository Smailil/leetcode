package main

/*
Дан массив целых чисел размера n.
Найдите все элементы, которые встречаются более ⌊ n/3 ⌋ раз.

Пример 1:
Вход: nums = [3,2,3]
Выход: [3]

Пример 2:
Вход: nums = [1]
Выход: [1]

Пример 3:
Вход: nums = [1,2]
Выход: [1,2]

Ограничения:
1 <= nums.length <= 5 * 10^4
-10^9 <= nums[i] <= 10^9

*/

// Алгоритм: модифицированный Бойер-Мур + верификация за проход.
func majorityElement(nums []int) []int {
	// Этап 1: отбор потенциальных кандидатов (не более 2)
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++

		if len(count) <= 2 {
			continue
		}

		// "Голосование": уменьшаем счётчики, удаляем нулевые
		newCount := make(map[int]int)
		for num, c := range count {
			if c > 1 {
				newCount[num] = c - 1
			}
		}
		count = newCount
	}

	// Этап 2: верификация кандидатов за проход по массиву
	// Вместо вложенного цикла считаем частоты всех кандидатов параллельно
	threshold := len(nums) / 3
	candidateFreq := make(map[int]int)

	for _, num := range nums {
		// Проверяем, является ли текущий элемент кандидатом
		if _, isCandidate := count[num]; isCandidate {
			candidateFreq[num]++
		}
	}

	// Формируем результат: только элементы, превысившие порог
	result := []int{}
	for candidate, freq := range candidateFreq {
		if freq > threshold {
			result = append(result, candidate)
		}
	}

	return result
}
