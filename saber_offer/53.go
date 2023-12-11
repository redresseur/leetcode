package main

/*
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109


注意：本题与主站 34 题相同（仅返回值不同）：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
	核心思路：找到第一个等于 target 的位置 left_target_index，
	和最后一个等于target 的位置 right_target_index,
	number = right_target_index - left_target_index + 1

	核心函数：search_target_boundary
	在二分查找函数的基础上进行改造：
	当查找到 nums[mid] == target 时，
	if 查找 left_target_index
		继续位移 high 指针，最终返回 low 指针
	else
		继续位移 low 指针，最终返回 high 指针
*/
func search_target_boundary(nums []int, target int, lower bool) (res int) {
	// 从左至右，找到第一个target的位置
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = low + (high-low)>>1
		if midNum := nums[mid]; midNum > target || (lower && midNum >= target) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if res = low; !lower {
		res = high
	}
	return
}

func search(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}
	// 从左至右，找到第一个target的位置
	low := search_target_boundary(nums, target, true)
	if low >= lenNums || nums[low] != target {
		// 如果找不到 target 则直接返回 0
		return 0
	}
	// 从左至右，找到最后一个target的位置
	// 从 [first, end] 区间检索
	nums = nums[low:]
	high := search_target_boundary(nums, target, false)
	if high < 0 || nums[high] != target {
		return 0
	}
	return high + 1
}
