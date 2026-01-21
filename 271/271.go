package main

import (
	"fmt"
	"strconv"
)

/*
Разработайте алгоритм для кодирования списка строк в одну строку. 
Затем закодированная строка декодируется обратно в исходный список строк.

Пример 1:
Input: ["neet","code","love","you"]
Output:["neet","code","love","you"]

Пример 2:
Input: ["we","say",":","yes"]
Output: ["we","say",":","yes"]

Ограничения:
0 <= strs.length < 100
0 <= strs[i].length < 200
strs[i] содержит только символы UTF-8.

*/

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, str := range strs {
		res += fmt.Sprintf("%d:%s", len(str), str)
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != ':' {
			j++
		}
		
		lengthStr := encoded[i:j]
		length, _ := strconv.Atoi(lengthStr)
		
		start := j + 1
		end := start + length
		str := encoded[start:end]
		
		result = append(result, str)
		
		i = end
	}
	return result
}