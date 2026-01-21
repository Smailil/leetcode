package main

/*
Спроектируйте стек, который поддерживает push, pop, top и получение минимального элемента за постоянное время.

Реализуйте класс MinStack:

MinStack() инициализирует объект стека.
void push(int val) помещает элемент val в стек.
void pop() удаляет элемент из вершины стека.
int top() получает верхний элемент стека.
int getMin() получает минимальный элемент в стеке.
Вы должны реализовать решение с временной сложностью O(1) для каждой функции.

*/

type MinStack struct {
    stack []int
	minStack []int 
}


func Constructor() MinStack {
    stack := make([]int, 0)
	minStack := make([]int, 0)
	return MinStack{stack, minStack}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
        this.minStack = append(this.minStack, val)
    } else {
		newMin := min(this.minStack[len(this.minStack)-1], val)
		this.minStack = append(this.minStack, newMin)
	}
}


func (this *MinStack) Pop()  {
    this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


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