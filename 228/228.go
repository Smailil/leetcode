package main

import "fmt"

func summaryRanges(nums []int) []string {
	out := make([]string, 0)
	if len(nums) == 0 {
		return out
	}
	begin := nums[0]
	end := nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}
		if end+1 == nums[i] {
			end++
		} else {
			if begin == end {
				out = append(out, fmt.Sprintf("%v", end))
			} else {
				out = append(out, fmt.Sprintf("%v->%v", begin, end))
			}
			begin = nums[i]
			end = nums[i]
		}
	}
	if begin == end {
		out = append(out, fmt.Sprintf("%v", end))
	} else {
		out = append(out, fmt.Sprintf("%v->%v", begin, end))
	}
	return out
}
