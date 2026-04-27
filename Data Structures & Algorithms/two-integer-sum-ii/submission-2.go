func twoSum(numbers []int, target int) []int {
	wantedNumMap := make(map[int]int) // maps the wanted number to its index
	for i := 0; i < len(numbers); i++ {
		want := target - numbers[i]
		index, exists := wantedNumMap[want]
		if exists {
			return []int{index + 1, i + 1}
		}
		wantedNumMap[numbers[i]] = i
	}

	return []int{}
}