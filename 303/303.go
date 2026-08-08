package main

/*
Дан массив целых чисел nums. Обработайте несколько запросов следующего типа:

Вычислите сумму элементов nums между индексами left и right включительно,
где left <= right.

Реализуйте класс NumArray:
NumArray(int[] nums) — инициализирует объект массивом целых чисел nums.
int sumRange(int left, int right) — возвращает сумму элементов nums между
индексами left и right включительно:
nums[left] + nums[left + 1] + ... + nums[right].

Пример 1:
Вход:
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
Выход:
[null, 1, -1, -3]
Пояснение:
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // возвращает (-2) + 0 + 3 = 1
numArray.sumRange(2, 5); // возвращает 3 + (-5) + 2 + (-1) = -1
numArray.sumRange(0, 5); // возвращает (-2) + 0 + 3 + (-5) + 2 + (-1) = -3

Ограничения:
1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
0 <= left <= right < nums.length
Будет выполнено не более 10^4 вызовов sumRange.
*/

// NumArray хранит префиксные суммы элементов исходного массива.
// Это переносит основную работу в Constructor и ускоряет каждый запрос.
type NumArray struct {
    prefix []int
}

// Constructor строит таблицу префиксных сумм для nums.
// После построения сумма диапазона вычисляется разностью двух значений.
// Время: O(n), где n = len(nums). Дополнительная память: O(n).
func Constructor(nums []int) NumArray {
    // prefix[i] хранит сумму nums[0] + ... + nums[i]. Фиктивный нулевой
    // элемент не используется, поэтому длина prefix совпадает с len(nums).
    prefix := make([]int, len(nums))

    // cur поддерживает сумму элементов от начала до текущего индекса i,
    // поэтому после присваивания prefix[i] содержит нужный префикс.
    cur := 0
    for i, num := range nums {
        cur += num
        prefix[i] = cur
    }
    return NumArray{prefix: prefix}
}

// SumRange возвращает сумму nums[left:right+1] по готовым префиксам.
// Время: O(1). Дополнительная память: O(1).
func (this *NumArray) SumRange(left int, right int) int {
    // prefix[right] содержит сумму от начала массива до right включительно.
    rightSum := this.prefix[right]

    // Для left == 0 вычитать нечего; проверка также не допускает обращения
    // к несуществующему элементу prefix[-1].
    leftSum := 0
    if left > 0 {
        leftSum = this.prefix[left-1]
    }

    // Вычитаем сумму до left, удаляя элементы вне требуемого диапазона.
    return rightSum - leftSum
}

/**
 * Объект NumArray будет создан и вызван следующим образом:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
