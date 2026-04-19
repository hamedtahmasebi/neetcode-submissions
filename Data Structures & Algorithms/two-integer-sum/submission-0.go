import (
	"slices"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				result := []int{i, j}
				slices.Sort(result)
				return result
			}
		}
	}
	return []int{-1, -1}
}