package main

func hashKey(s string) [26]byte {
	key := [26]byte{}
	for i := range s {
		key[s[i]-'a']++
	}
	return key

}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[[26]byte][]string)
	for _, str := range strs {
		key := hashKey(str)
		strMap[key] = append(strMap[key], str)

	}
	res := make([][]string, 0, len(strMap))
	for _, val := range strMap {
		res = append(res, val)
	}
	return res
}
