package main

import "sort"

/*
Дан целочисленный массив nums, верните все тройки [nums[i], nums[j], nums[k]] такие, что i != j, i != k и j != k,
а также nums[i] + nums[j] + nums[k] == 0.
Обратите внимание, что набор решений не должен содержать повторяющихся троек.

Пример 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Объяснение:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
Отдельные тройки — это [-1,0,1] и [-1,-1,2].
Обратите внимание, что порядок вывода и порядок троек не имеют значения.

Пример 2:
Input: nums = [0,1,1]
Output: []
Объяснение: Единственная возможная тройка не дает в сумме 0.

Пример 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Объяснение: Сумма единственно возможной тройки равна 0.

*/

func threeSum(nums []int) [][]int {
    if len(nums) < 3 {
        return [][]int{}
    }
    
    sort.Ints(nums)
    result := make([][]int, 0)
    
    for i := 0; i < len(nums)-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
		if nums[i] > 0 {
            break
        }
        
        l, r := i+1, len(nums)-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum < 0 {
                l++
            } else if sum > 0 {
                r--
            } else {
                result = append(result, []int{nums[i], nums[l], nums[r]})

                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }
                l++
                r--
            }
        }
    }
    
    return result
}
