package main

/*
剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

解题思路
参考官方方案使用双指针方法，进行了一点优化：
将刷新窗口 windows map 的操作从逐个移除变为判断 prev < left，如果满足该条件说明是历史未清理的数据，按照未存在处理。
https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
*/
func lengthOfLongestSubstring(s string) int {
	win, max, left := map[byte]int{}, 0, 0
	for i := 0; i < len(s); i++ {
		prev, ok := win[s[i]]
		win[s[i]] = i

		if !ok || prev < left { // 正常累加
			if max < i-left+1 {
				max = i - left + 1
			}
		} else { // 遇到重复，位移 left
			if max < i-left {
				max = i - left
			}
			left = prev + 1
		}
	}
	return max
}
