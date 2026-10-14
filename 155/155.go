package main

/*
Спроектируйте стек, который поддерживает push, pop, top
и получение минимального элемента за постоянное время.

Реализуйте класс MinStack:
MinStack() инициализирует объект стека.
void push(int val) помещает элемент val в стек.
void pop() удаляет элемент из вершины стека.
int top() получает верхний элемент стека.
int getMin() получает минимальный элемент в стеке.
Вы должны реализовать решение с временной сложностью O(1) для каждой функции.
*/

// MinStack хранит значения и минимумы соответствующих префиксов.
// minStack[i] — минимум stack[:i+1]; длины двух срезов всегда совпадают.
// Буферы O(h), h — максимальная достигнутая длина, а не текущая:
// Pop не освобождает общие массивы при сокращении длины.
type MinStack struct {
    stack []int
	minStack []int 
}


// Constructor создаёт пустой стек и пустую историю минимумов.
// Время, рабочая память и начальный результат O(1).
func Constructor() MinStack {
    stack := make([]int, 0)
	minStack := make([]int, 0)
	return MinStack{stack, minStack}
}


// Push добавляет значение и новый минимум префикса, меняя оба поля this.
// Время O(1) амортизированно, O(h) при расширении.
// h — максимальная достигнутая глубина стека, как и для ёмкости буферов.
// Дополнительные буферы при расширении O(h); итоговое состояние O(h).
func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
	// Первое значение само минимально; иначе сравниваем с прежним минимумом.
	if len(this.minStack) == 0 {
        this.minStack = append(this.minStack, val)
    } else {
		newMin := min(this.minStack[len(this.minStack)-1], val)
		this.minStack = append(this.minStack, newMin)
	}
}


// Pop убирает вершину и её минимум, меняя this и сохраняя ёмкость массивов.
// Время и рабочая память O(1), результата нет.
// Требует непустой стек, иначе срез с длиной -1 вызовет панику.
func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


// Top читает вершину без изменения стека за O(1) времени и рабочей памяти.
// Результат int занимает O(1); пустой стек вызовет панику индекса.
func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


// GetMin берёт минимум всего стека из вершины истории минимумов.
// Время, рабочая память и результат O(1); this не изменяется.
// Требует непустой стек, иначе отрицательный индекс вызовет панику.
func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}


/**
 * Ваш объект MinStack будет создан и вызван следующим образом:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */