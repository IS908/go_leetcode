package leetcode

import "sort"

/**
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。

示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]

提示:
1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/
func combinationSum2(candidates []int, target int) [][]int {
	var result = make([][]int, 0)
	if len(candidates) == 0 {
		return result
	}
	sort.Ints(candidates)
	combinSum2(candidates, []int{}, 0, target, &result)
	return result
}

func combinSum2(candidates []int, list []int, idx, remain int, result *[][]int) {
	if remain < 0 {
		return
	}
	if remain == 0 {
		var tmp []int
		tmp = append(tmp, list...)
		*result = append(*result, tmp)
		return
	}

	for i := idx; i < len(candidates); i++ {
		var tmp = append(list, candidates[i])
		combinSum2(candidates, tmp, i+1, remain-candidates[i], result)
		for ; i < len(candidates)-1 && candidates[i] == candidates[i+1]; i++ {
		}
	}
}
