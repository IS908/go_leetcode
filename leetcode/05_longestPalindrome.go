/**
 * @Author: KevinCK
 * @Date: 2021/12/23 22:46
 */

package leetcode

func longestPalindrome(s string) string {
	var subStr = s[0:1]
	for i := 0; i < len(s); i++ {
		for step := 1; i-step >= 0 && i+step < len(s); step++ {
			if s[i-step] != s[i+step] {
				break
			}
			if len(subStr) < 1+2*step {
				subStr = s[i-step : i+step+1]
			}
		}
	}
	for i := 0; i < len(s); i++ {
		for step := 0; i-step >= 0 && i+1+step < len(s); step++ {
			if s[i-step] != s[i+1+step] {
				break
			}
			if len(subStr) < 2+2*step {
				subStr = s[i-step : i+1+step+1]
			}
		}
	}
	return subStr
}
