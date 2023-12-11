package bytedance

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：
输入：s = "cbbd"
输出："bb"
提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/

/*
	抽象成一个函数，好处就是可以同时处理 奇数回文和偶数回文的情况
	所有存在奇数偶数的情况都可以用这个方法来解决
*/
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 此处一定要加上，否则有可能出现panic，
	//	在最后一个指针的奇数情况时候，会遇到这个问题
	if l+1 > r {
		return ""
	}
	// 为什么是 l+1?
	// 终止有三种情况：l 越界，r 越界，不相等
	// 无论是哪一种，l 都多减了一个 1
	return s[l+1 : r]
}

func longestPalindrome(s string) (max string) {
	for i := range s {
		if sub := palindrome(s, i, i); len(sub) > len(max) {
			max = sub
		}

		if sub := palindrome(s, i, i+1); len(sub) > len(max) {
			max = sub
		}
	}
	return
}
