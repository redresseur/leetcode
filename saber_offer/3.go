package main

/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3


限制：

2 <= n <= 100000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
	解题思路：要有一个统计器来统计每一个数的状态，可以用 hash-map、bit-map
	此题比较特殊的地方在于容量为 n 的数组中 数的范围 是 [0,n)，那也就是说可以
	**用数组本身作为统计器**:
		nums[num] == num : 已存在
		nums[num] != num : 不存在
	从左向右遍历：

	for i := 0
		// check 当前状态
		if nums[i] == i // 此处已经被标记占用，跳过
			i++
			continue
		// 检查 num = nums[i] 的统计状态
		if num == nums[num] // 说明已存在，直接返回 num
			return num
		// 标记为已存在:
		swap nums[num], nums[i] // nums 的元素 i 尚未被标记占用，可以用来做临时寄存，

	return -1 // 没有重复的 num
*/
func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == i {
			i++
			continue
		}

		num := nums[i]
		if nums[num] == num {
			return num
		}

		nums[i], nums[num] = nums[num], nums[i]
	}

	return nums[0]
}
