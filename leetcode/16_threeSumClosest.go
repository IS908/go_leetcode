package leetcode

import (
	"sort"
)

/**
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。

示例 1：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：
输入：nums = [0,0,0], target = 1
输出：0

提示：
3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res = nums[0] + nums[1] + nums[2]

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < len(nums)-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third := len(nums) - 1; third > second; third-- {
				if third < len(nums)-1 && nums[third] == nums[third+1] {
					continue
				}
				sum := nums[first] + nums[second] + nums[third]
				if abs(sum-target) < abs(res-target) {
					res = sum
				}
			}

		}
	}
	return res
}

func threeSumClosest1(nums []int, target int) int {
	sort.Ints(nums)
	var res = nums[0] + nums[1] + nums[2]

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second, third := first+1, len(nums)-1
		for second < third {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if third < len(nums)-1 && nums[third] == nums[third+1] {
				third--
				continue
			}
			sum := nums[first] + nums[second] + nums[third]
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				second++
			} else {
				third--
			}
		}
	}
	return res
}
