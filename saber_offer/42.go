package main

/*
剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。



示例1:

输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。


提示：

1 <= arr.length <= 10^5
-100 <= arr[i] <= 100
注意：本题与主站 53 题相同：https://leetcode-cn.com/problems/maximum-subarray/
*/

// 标签：动态规划 剑指 Offer 42. 连续子数组的最大和
/*
解题思路
状态转移方程：cc = max(cc + nums[i], nums[i]), cc 为子数组累加值；

max: 记录最大的 cc；

时间复杂度：O(N) 空间复杂度O(1)

作者：luckydogchina
链接：https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/by-luckydogchina-9q15/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	max, cc := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if cc = nums[i] + cc; cc < nums[i] {
			cc = nums[i]
		}

		if cc > max {
			max = cc
		}
	}
	return max
}
