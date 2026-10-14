package main
/*
Дан целочисленный массив nums и целое число k.
Верните k наиболее часто встречающихся элементов.
Вы можете вернуть ответ в любом порядке.

Пример 1:
Вход: nums = [1,1,1,2,2,3], k = 2
Выход: [1,2]

Пример 2:
Вход: nums = [1], k = 1
Выход: [1]
*/
// topKFrequent группирует числа по частоте и выбирает k из самых частых.
// Среднее время O(n+d), n=len(nums), d — число разных значений; это O(n).
// Рабочая память O(n+d): карта частот и корзины; результат O(k).
// Вход не меняется; требуется 1 <= k <= d. Равные частоты идут в порядке map.
func topKFrequent(nums []int, k int) []int {
    hm := make(map[int]int)
	for _, num := range nums {
		hm[num]++
	}
	// Частота лежит в [1,n]; корзина 0 не используется.
	buckets := make([][]int, len(nums) + 1)
	for num, count := range hm {
		buckets[count] = append(buckets[count], num)
	}
	resultArray := make([]int, 0, k)
	// Идём от самой высокой частоты вниз; каждый ключ появляется в одной
	// корзине, а остановка после k элементов даёт требуемый размер ответа.
	for i := len(nums); i != 0; i-- {
		for _, num := range buckets[i] {
			resultArray = append(resultArray, num)
			if len(resultArray) == k {
				return resultArray
			}
		}
	}
	return resultArray
}