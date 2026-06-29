package main

/*
Given an integer array arr, return the length of a maximum size turbulent subarray of arr.

A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:

For i <= k < j:
arr[k] > arr[k + 1] when k is odd, and
arr[k] < arr[k + 1] when k is even.
Or, for i <= k < j:
arr[k] > arr[k + 1] when k is even, and
arr[k] < arr[k + 1] when k is odd.
 

Example 1:

Input: arr = [9,4,2,10,7,8,8,1,9]
Output: 5
Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]
Example 2:

Input: arr = [4,8,12,16]
Output: 2
Example 3:

Input: arr = [100]
Output: 1
 

Constraints:

1 <= arr.length <= 4 * 104
0 <= arr[i] <= 109

*/

func maxTurbulenceSize(arr []int) int {
    // l и r — левая и правая границы скользящего окна.
    // res — максимальная найденная длина подмассива (минимум 1, так как один элемент — это уже подмассив).
    l, r, res := 0, 1, 1
    
    // prev сохраняет знак предыдущего сравнения соседних элементов: 
    // ">" (убывание) или "<" (возрастание). Это нужно для проверки чередования знаков.
    prev := ""

    // Итерируемся по массиву, пока правый указатель не выйдет за его пределы
    for r < len(arr) {
        // Случай 1: текущая пара элементов убывает (arr[r-1] > arr[r]), 
        // и предыдущее сравнение НЕ было убыванием (prev != ">")
        if arr[r-1] > arr[r] && prev != ">" {
            // Условие чередования выполнено — обновляем максимум длины текущего окна
            if r-l+1 > res {
                res = r - l + 1
            }
            r++          // Сдвигаем правую границу, расширяя окно
            prev = ">"   // Фиксируем текущий знак сравнения для следующей итерации

        // Случай 2: текущая пара элементов возрастает (arr[r-1] < arr[r]), 
        // и предыдущее сравнение НЕ было возрастанием (prev != "<")
        } else if arr[r-1] < arr[r] && prev != "<" {
            // Условие чередования выполнено — обновляем максимум длины текущего окна
            if r-l+1 > res {
                res = r - l + 1
            }
            r++          // Сдвигаем правую границу
            prev = "<"   // Фиксируем текущий знак сравнения

        // Случай 3: условие турбулентности нарушено (знаки совпали или элементы равны)
        } else {
            // Если соседние элементы равны, они не могут быть частью турбулентного подмассива.
            // Пропускаем этот элемент, сдвигая правый указатель.
            if arr[r] == arr[r-1] {
                r++
            }
            
            // Сбрасываем левую границу окна. 
            // Новое потенциальное окно начинаем строить с элемента (r - 1).
            l = r - 1
            
            // Сбрасываем сохраненный знак, так как новое окно только начинается
            prev = ""
        }
    }

    return res
}