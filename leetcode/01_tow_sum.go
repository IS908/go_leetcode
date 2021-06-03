package leetcode

func twoSum(nums []int, target int) []int {
	for index, num := range nums {
		for i := index + 1; i < len(nums); i++ {
			if num+nums[i] == target {
				return []int{index, i}
			}
		}
	}
	return nil
}

func twoSum_01(nums []int, target int) []int {
	var (
		gap    int
		numMap = make(map[int]int, 0)
	)
	for index, num := range nums {
		numMap[num] = index
	}

	for index, num := range nums {
		gap = target - num
		if i, ok := numMap[gap]; ok && index != i {
			return []int{index, i}
		}
	}
	return nil
}
