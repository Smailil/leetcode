package main

/*
Дан вложенный список целых чисел nestedList.
Каждый элемент — целое число или список, который тоже может содержать целые
числа и другие списки.
Реализуйте итератор, который обходит этот список как плоскую последовательность.

Реализуйте класс NestedIterator:
NestedIterator(List<NestedInteger> nestedList)
    Инициализирует итератор вложенным списком nestedList.
int next() — возвращает следующее целое число во вложенном списке.
boolean hasNext() — возвращает true, если во вложенном списке ещё есть
целые числа, и false в противном случае.

Ваш код проверяется следующим псевдокодом:
инициализировать iterator списком nestedList
res = []
пока iterator.hasNext()
    добавить iterator.next() в конец res
вернуть res
Решение считается правильным, если res совпадает с ожидаемым плоским списком.

Пример 1:
Вход: nestedList = [[1,1],2,[1,1]]
Выход: [1,1,2,1,1]
Пояснение: Вызываем next, пока hasNext не вернёт false.
Значения, которые возвращает next, идут в порядке [1,1,2,1,1].

Пример 2:
Вход: nestedList = [1,[4,[6]]]
Выход: [1,4,6]
Пояснение: Вызываем next, пока hasNext не вернёт false.
Значения, которые возвращает next, идут в порядке [1,4,6].

Ограничения:
1 <= nestedList.length <= 500
Целые числа во вложенном списке лежат в диапазоне [-10^6, 10^6].
*/

/**
 * // Этот интерфейс позволяет создавать вложенные списки.
 * // Не реализуйте его и не делайте предположений о его реализации.
 * type NestedInteger struct {
 * }
 *
 * // Возвращает true, если NestedInteger хранит целое число, а не список.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Возвращает целое число, которое хранит NestedInteger.
 * // Если хранится вложенный список, результат не определён.
 * // Перед вызовом нужно проверить, что хранится целое число.
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Задаёт целое число, которое хранит NestedInteger.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Задаёт тип содержимого «вложенный список» и добавляет в него elem.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Возвращает вложенный список, который хранит NestedInteger.
 * // Длина списка равна нулю, если NestedInteger хранит целое число.
 * // Обращайтесь к полю List напрямую, если хотите изменить список.
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// NestedIterator обходит числа слева направо, раскрывая списки по мере нужды.
type NestedIterator struct {
	// На вершине — следующий узел для обхода: число или вложенный список.
	stack []*NestedInteger
}

// Constructor подготавливает обход, не меняя порядок в исходном nestedList.
// Время: O(m), где m = len(nestedList); дополнительная память: O(m).
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// Копируем только ссылки верхнего уровня, чтобы развернуть свою копию.
	stack := make([]*NestedInteger, len(nestedList))
	copy(stack, nestedList)
	// Левый элемент должен оказаться в конце среза, на вершине стека.
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return &NestedIterator{stack: stack}
}

// Next снимает подготовленное число с вершины стека.
// Вызывать после HasNext, вернувшего true, пока число не было снято.
// Время: O(1). Дополнительная память: O(1).
func (it *NestedIterator) Next() int {
	// HasNext гарантирует, что стек непуст и GetInteger вызван для числа.
	val := it.stack[len(it.stack)-1].GetInteger()
	it.stack = it.stack[:len(it.stack)-1]
	return val
}

// HasNext раскрывает списки на вершине, пока не найдёт следующее число.
// Время: до O(n) за вызов, где n — число узлов (чисел и списков).
// За q вызовов время O(n + q): каждый список раскрывается не более раза,
// а остановка на числе или на пустом стеке занимает O(1).
// Дополнительная память итератора: O(n) в худшем случае.
func (it *NestedIterator) HasNext() bool {
	for len(it.stack) > 0 {
		top := it.stack[len(it.stack)-1]
		if top.IsInteger() {
			// Оставляем число для Next: повторный HasNext его не снимет.
			return true
		}
		// Вместо списка кладём его детей; пустой список ничего не добавит.
		it.stack = it.stack[:len(it.stack)-1]
		nestedList := top.GetList()
		// Кладём детей справа налево, чтобы первым обработать левого.
		// Для [4,[6]] сначала кладём [6], затем 4: на вершине будет 4.
		for i := len(nestedList) - 1; i >= 0; i-- {
			it.stack = append(it.stack, nestedList[i])
		}
	}
	return false
}