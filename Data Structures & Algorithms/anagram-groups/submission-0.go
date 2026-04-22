func hashStr(s string) string {
	// hashes the string into it's characters + count of characters sorted decending in ASCII order, e.g. cat -> t1c1a1
	var charCounts [26]int
	for _, char := range s {
		charCounts['z'-char]++
	}
	resultStrBuilder := strings.Builder{}
	resultStrBuilder.Grow(len(s) * 2)
	for i, count := range charCounts {
		if count == 0 {
			continue
		}
		resultStrBuilder.WriteRune(rune('z' - i))
		resultStrBuilder.WriteString(strconv.Itoa(int(count)))
	}

	return resultStrBuilder.String()
}

func groupAnagrams(strs []string) [][]string {
	var hashmap = make(map[string][]string) // string hash to strings in the original array

	for idx, str := range strs {
		h := hashStr(str)
		_, ok := hashmap[h]
		if !ok {
			hashmap[h] = []string{strs[idx]}
			continue
		}
		hashmap[h] = append(hashmap[h], strs[idx])
	}

	result := [][]string{}
	for _, strs := range hashmap {
		result = append(result, strs)
	}

	return result
}