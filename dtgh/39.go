package dtgh

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
链接：https://leetcode-cn.com/problems/combination-sum
*/

// dfs 组合法
// 参考：https://leetcode-cn.com/problems/combination-sum/solution/jia-fa-de-dfs-by-jin-ai-yi/
func combinationSum1(candidates []int, target int) [][]int {
	var res = [][]int{}
	combination(candidates, target, 0, []int{}, &res)
	return res
}

func combination(candidates []int, target int, pathSum int, path []int, res *[][]int) {
	if pathSum >= target {
		return
	}

	for i, c := range candidates {
		tmp := append(path, c)
		if pathSum+c == target {
			answer := make([]int, len(tmp))
			copy(answer, tmp)
			*res = append(*res, answer)
		}

		combination(candidates[i:], target, pathSum+c, tmp, res)
	}
}

// dfs + 剪枝法，
// 参考：https://leetcode-cn.com/problems/combination-sum/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-2/
func combinationSum(candidates []int, target int) [][]int {
	visited := make([]int, target+1)
	dp := make([][]int, target+1)
	visited[0] = 1
	for i := 0; i <= target; i++ {
		if visited[i] != 1 {
			continue
		}

		for _, c := range candidates {
			prv := i + c
			if prv > target {
				continue
			} else if prv <= target {
				if len(dp[prv]) <= 0 {
					dp[prv] = make([]int, target+1)
				}
				visited[prv] = 1
				dp[prv][i] = c
			}
		}
	}

	res := output(dp, 0)
	return res
}

func output(dp [][]int, pprv int) [][]int {
	top := len(dp) - 1
	if top == 0 {
		return [][]int{{}}
	}

	res := [][]int{}
	for i, prv := range dp[top] {
		if prv == 0 {
			continue
		}
		if prv < pprv {
			break
		}
		tmp := output(dp[0 : i+1], prv)
		for _, t := range tmp {
			t = append(t, prv)
			res = append(res, t)
		}
	}

	return res
}