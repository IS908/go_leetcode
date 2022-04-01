package leetcode

/**
 * 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3
示例 2：
输入：nums = [3,4,-1,1]
输出：2
示例 3：
输入：nums = [7,8,9,11,12]
输出：1

提示：
1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1
*/
func firstMissingPositive(nums []int) int {
	var n = len(nums)
	for i := 0; i < n; i++ {
		tmp := nums[i]
		if tmp < 0 || tmp > n {
			continue
		} else {
			for tmp > 0 && tmp <= n {
				t := nums[tmp-1]
				nums[tmp-1] = tmp
				if tmp == t {
					break
				}
				tmp = t
			}
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
