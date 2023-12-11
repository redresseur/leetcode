package main

/*
剑指 Offer 53 - II. 0～n-1中缺失的数字
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
示例 1:
输入: [0,1,3]
输出: 2
示例 2:
输入: [0,1,2,3,4,5,6,7,9]
输出: 8
限制：
1 <= 数组长度 <= 10000
通过次数256,494提交次数573,545
*/

/*
	核心思路：参考了剑指 offer 第 11 题（在一个旋转数组中查找最小值），
	在一个”有序数组“中查找某个特异点，一般左右指针移动的时候，
	只有一个需要 mid+1（11 中是 high） 或 mid -1（本题中是 low），
	另一个继承了（mid），这完全取决于 mid 元素 是否可能为解。
	比如等值二分查找中，mid如果不等，则不可能为解，直接”抛弃“即可。

	而本题中，当 nums[mid] > mid 时，mid就有可能是解，
	而在 nums[mid] == mid 时，不可能为解。
	另外总结一个一般规律，当二分查找指定一个 target（等值查找、等值范围查找） 时，
	循环终止条件为 low<=high，而不指定 target 时，循环终止条件为 low < high.
*/
func missingNumber(nums []int) int {
	low, high, mid := 0, len(nums), 0
	for low < high { // [left, right)
		mid = low + (high-low)>>1
		if nums[mid] > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}
