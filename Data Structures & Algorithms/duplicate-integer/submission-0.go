func hasDuplicate(nums []int) bool {
	seenMap := make(map[int]bool, len(nums))
	for _, n := range nums {
		if seenMap[n] {
			return true
		}
		seenMap[n] = true
	}
	return false
}