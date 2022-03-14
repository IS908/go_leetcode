package leetcode

import "math"

/**
 * 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
 * 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
 * 假设环境不允许存储 64 位整数（有符号或无符号）。
 */
func reverse(x int) int {
	if x == 0 {
		return x
	}
	var (
		ret      int
		nagetive bool
	)
	if x < 0 {
		nagetive = true
		x = -x
	}
	for x > 0 {
		if ret > math.MaxInt32/10 {
			return 0
		}
		var n = x % 10
		x /= 10
		ret = ret*10 + n
	}
	if nagetive {
		return -ret
	}
	return ret
}
