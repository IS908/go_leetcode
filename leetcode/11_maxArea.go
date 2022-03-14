package leetcode

/**
 * 给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。
*/
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var max = 0
	for i, j := 0, len(height)-1; i < j; {
		var minH = height[i]
		if minH > height[j] {
			minH = height[j]
			if tmp := minH * (j - i); tmp > max {
				max = tmp
			}
			j--
		} else {
			if tmp := minH * (j - i); tmp > max {
				max = tmp
			}
			i++
		}
	}
	return max
}
