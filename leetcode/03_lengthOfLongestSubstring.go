package leetcode

/**
 * @Author: KevinCK
 * @Date: 2021/12/23 00:14
 */

func lengthOfLongestSubstring(s string) int {
	var (
		bitMap = make(map[int32]int)
		start  int
		curLen int
		maxLen int
	)

	for i, b := range s {
		if index, ok := bitMap[b]; ok {
			if index > start {
				start = index
			}
			curLen = i - start
		} else {
			curLen++
		}
		if curLen > maxLen {
			maxLen = curLen
		}
		bitMap[b] = i
	}

	return maxLen
}
