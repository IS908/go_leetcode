/**
 * @Author: KevinCK
 * @Date: 2021/12/23 21:42
 */

package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l = len(nums1) + len(nums2)
	var nums = make([]int, 0, l)
	var i, i1, i2 = 0, 0, 0
	for ; i1 < len(nums1) && i2 < len(nums2); i++ {
		if nums1[i1] < nums2[i2] {
			nums[i] = nums1[i1]
			i1++
		} else {
			nums[i] = nums2[i2]
			i2++
		}
	}

	for ; i < l; i++ {
		if i1 < len(nums1) {
			nums[i] = nums1[i1]
			i1++
		} else if i2 < len(nums2) {
			nums[i] = nums2[i2]
			i2++
		}
	}

	if l == 1 {
		return float64(nums[0])
	} else if l&1 == 0 {
		l = l >> 1
		return float64(nums[l-1]+nums[l]) / 2
	} else {
		return float64(nums[l>>1])
	}
}
