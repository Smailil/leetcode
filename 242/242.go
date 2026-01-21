package main

func isAnagram(s string, t string) bool {
	chars := make([]int, 26)
	for _, c := range s {
		i := int(c - 'a')
		chars[i]++
	}
	for _, c := range t {
		i := int(c - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}
