/**
 * @Author: KevinCK
 * @Date: 2021/12/27 23:15
 */

package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var (
		btsArr = make([]string, numRows)
		numIdx int
		flag   = -1
		res    string
	)

	for _, b := range s {
		btsArr[numIdx] += string(b)
		if numIdx == 0 || numIdx == numRows-1 {
			flag = -flag
		}
		numIdx += flag
	}

	for _, arr := range btsArr {
		res += arr
	}
	return res
}
