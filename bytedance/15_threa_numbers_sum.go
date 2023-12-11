package bytedance

import "sort"

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]
提示：
0 <= nums.length <= 3000
-105 <= nums[i] <= 105

https://leetcode.cn/problems/3sum/
*/

// 将三数之和降维为两数之和
func findTarget(nums []int, target int) (res [][]int) {
	for l, r := 0, len(nums)-1; l < r; {
		if l > 0 && nums[l] == nums[l-1] { // 跳过重复的
			l++
			continue
		}
		if sum := nums[r] + nums[l]; sum == target {
			res = append(res, []int{-target, nums[l], nums[r]})
			r--
			l++
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return res
}

func threeSum(nums []int) (res [][]int) {
	if len(nums) <= 2 {
		return
	}
	sort.Slice(nums, func(i, j int) bool { // 排序
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的 target
			continue
		}

		res = append(res, findTarget(nums[i+1:], -nums[i])...)
	}
	return
}
