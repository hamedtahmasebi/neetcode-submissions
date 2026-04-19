func isAnagram(s string, t string) bool {
	seenCharsMap := make(map[rune]*[2]uint, len(s))
	for _, char := range s {
		if seenCharsMap[char] == nil {
			seenCharsMap[char] = &[2]uint{0, 0}
		}
		seenCharsMap[char][0] += 1
	}

	for _, char := range t {
		if seenCharsMap[char] == nil || seenCharsMap[char][0] == 0 {
			return false
		}
		seenCharsMap[char][1] += 1
	}

	for _, arr_ptr := range seenCharsMap {
		if arr_ptr[0] != arr_ptr[1] {
			return false
		}
	}

	return true
}
