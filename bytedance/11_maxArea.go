package bytedance

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
https://leetcode.cn/problems/container-with-most-water/
核心思路： 双指针 盛水最多水的容器
if
 	min(left, right) == left
then
	min(left, right - 1) <= left
	(right - left - 1)*min(left, right - 1) < (right*left)*min(left, right)
此时只有 left 节点右移才有可能找到更大的容器，min(left, right) == right 时同理，
right 节点左移才能找到更大的容器。
*/
func maxArea(height []int) (max int) {
	left, right := 0, len(height)-1
	for left < right {
		width, min := right-left, height[left]
		if min > height[right] {
			min = height[right]
			right--
		} else {
			left++
		}

		if cap := min * width; cap > max {
			max = cap
		}
	}
	return
}
