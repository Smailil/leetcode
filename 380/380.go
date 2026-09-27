package main

import "math/rand"

/*
Реализуйте класс RandomizedSet:

RandomizedSet() инициализирует объект RandomizedSet.
bool insert(int val) добавляет val в множество, если его там нет.
Возвращает true, если val отсутствовал, и false в противном случае.
bool remove(int val) удаляет val из множества, если он там есть.
Возвращает true, если val присутствовал, и false в противном случае.
int getRandom() возвращает случайный элемент текущего множества.
Гарантируется, что при вызове метода множество не пусто.
Каждый элемент должен возвращаться с одинаковой вероятностью.
Каждый метод должен работать в среднем за O(1).

Пример 1:
Вход:
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Выход:
[null, true, false, true, 2, true, false, 2]
Пояснение:
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Добавляет 1 и возвращает true.
randomizedSet.remove(2); // Возвращает false: числа 2 нет в множестве.
randomizedSet.insert(2); // Добавляет 2 и возвращает true.
                         // Множество теперь содержит [1,2].
randomizedSet.getRandom(); // Случайным образом возвращает 1 или 2.
randomizedSet.remove(1); // Удаляет 1 и возвращает true.
                         // Множество теперь содержит [2].
randomizedSet.insert(2); // Число 2 уже есть, поэтому возвращает false.
randomizedSet.getRandom(); // В множестве только 2, поэтому вернёт 2.

Ограничения:
-2^31 <= val <= 2^31 - 1
Не более 2 * 10^5 вызовов методов insert, remove и getRandom.
При вызове getRandom в структуре данных есть хотя бы один элемент.

*/

// RandomizedSet хранит уникальные значения в плотном срезе nums.
// numMap[val] всегда равен индексу val в nums: карта даёт быстрый поиск,
// а плотный срез позволяет выбирать любой элемент по случайному индексу.
// Память O(k), где k — максимальный размер множества за время его работы:
// выделенная ёмкость может оставаться после удаления элементов.
type RandomizedSet struct {
    numMap map[int]int
    nums   []int
}

// Constructor создаёт пустое множество за O(1) времени и памяти.
func Constructor() RandomizedSet {
    return RandomizedSet{
        numMap: make(map[int]int),
        nums:   []int{},
    }
}

// Insert добавляет val, только если его ещё нет, и сообщает об успехе.
// Время в среднем амортизированно O(1); хранение каждого нового элемента
// требует O(1) памяти амортизированно, но расширение может стоить O(k).
func (this *RandomizedSet) Insert(val int) bool {
    // Карта сразу определяет, есть ли val; повторная вставка ничего не меняет.
    if _, exists := this.numMap[val]; exists {
        return false
    }
    // Новый элемент попадёт в конец nums: запоминаем его индекс до append.
    this.numMap[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

// Remove удаляет val, если он есть, сохраняя nums без пустых позиций.
// Время в среднем O(1) благодаря карте; рабочая память O(1).
func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.numMap[val]; !exists {
        return false
    }
    // Заменяем val последним элементом, чтобы не сдвигать хвост среза.
    // Его новый индекс нужно записать в карте до удаления val.
    idx := this.numMap[val]
    last := this.nums[len(this.nums)-1]
    this.nums[idx] = last
    this.numMap[last] = idx
    // Если val был последним, обновлённая запись тут же удалится.
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.numMap, val)
    return true
}

// GetRandom возвращает равновероятно выбранное значение из nums.
// Время в среднем O(1), рабочая память O(1); по условию nums не пуст.
func (this *RandomizedSet) GetRandom() int {
    // Каждый индекс одинаково вероятен, а значения в nums не повторяются.
    return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Объект RandomizedSet будет создан и использован так:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
