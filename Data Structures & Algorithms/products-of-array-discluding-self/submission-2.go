func productExceptSelf(nums []int) []int {
	// divide and conquer baby
	length := len(nums)
	prefix := make([]int, length)
	suffix := make([]int, length)
	results := make([]int, length)

	prefix[0] = 1
	suffix[length-1] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := length - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := range results {
		results[i] = prefix[i] * suffix[i]
	}

	return results
}