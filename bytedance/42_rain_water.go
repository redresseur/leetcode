package bytedance

/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

https://leetcode.cn/problems/trapping-rain-water/
*/
// 核心思路：单调栈 双指针 动态规划
func trap(height []int) (water int) {
	if len(height) <= 2 {
		return 0
	}

	stack := []int{}
	for i := 0; i < len(height); i++ {
		for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			left, low := stack[len(stack)-1], height[i]
			if low > height[left] {
				low = height[left]
			}

			water += (i - left - 1) * (low - height[top])
		}
		stack = append(stack, i)
	}
	return
}

func trapDoublePointer(height []int) (water int) {
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if leftMax < height[left] {
			leftMax = height[left]
		}
		if rightMax < height[right] {
			rightMax = height[right]
		}
		if height[left] < height[right] {
			water += leftMax - height[left]
			left++
		} else {
			water += rightMax - height[right]
			right--
		}
	}
	return
}
