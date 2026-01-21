package main

import "fmt"

func compress(chars []byte) int {
	i, res := 0, 0
	n := len(chars)
	for i < n {
		c := chars[i]
		count := 0
		for i < n && chars[i] == c {
			count++
			i++
		}
		if count == 1 {
			chars[res] = c
			res++
		} else {
			in := fmt.Sprintf("%c%d", c, count)
			for _, char := range []byte(in) {
				chars[res] = char
				res++
			}
		}
	}
	return res
}
