package bytedance

/*
31. 下一个排列
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

https://leetcode.cn/problems/next-permutation/
*/

/*
解题思路
解题思路实际是用了分治(分割为子问题)的思想，以尾部tail为起点，向前扩充：
	pivot = tail - 1
	while nums(pivot,tail) is max:
		pivot--

为什么 max 要跳过？
	因为max下一个状态就是min，如果当前整体 nums 不为 max，进入 min 显然不合理。

如何判断为 max 状态？
	nums(pivot,tail) 为降序。

找到 nums(pivot, tail) 不是 max 组合时，将它转化为下一个状态：
	nums[pivot] 和 nums(pivot+1,tail) 中大于自己的最小值 min(>nums[pivot]) 交换，交换后显然大于原组合，当并不是下一状态，因为此时 nums(pivot+1,tail) 是 max 状态，需要转化为 min 状态。

时间复杂度： O(N^2)
空间复杂度： O(1)

作者：luckydogchina
链接：https://leetcode.cn/problems/next-permutation/solution/liang-ci-bian-li-guan-fang-da-an-by-luck-5ip5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func nextPermutation(nums []int) {
	// 查找转折点
	pivot, tail := -1, len(nums)-1
	for i := tail; i > 0; i-- {
		if nums[i] > nums[i-1] {
			pivot = i - 1 // 此处注意，转折点是 i - 1！！！
			break
		}
	}

	if pivot >= 0 {
		// 找到大于nums[pivot] 中最小的值
		for i := tail; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				nums[i], nums[pivot] = nums[pivot], nums[i] // 交换值
				break
			}
		}
	}

	// 反转数组
	l, r := pivot+1, tail
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return
}
