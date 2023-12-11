package main

/*剑指 Offer 46. 把数字翻译成字符串
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
示例 1:

输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

提示：

0 <= num < 231
解题思路: 从前往后还是从后往前都是一样的，以 10086 为例：

round1: 6 86 1
round2: 8 08 1
round3: 0 00 1
round4: 0 10 2
round5: 1 1  2

作者：luckydogchina
链接：https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/solution/by-luckydogchina-fdr1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func translateNum(num int) int {
	pre, pre1 := 1, 1
	for n := num; n != 0; n /= 10 {
		bpi := pre
		if n%100 < 26 && n%100 != n%10 {
			bpi += pre1
		}
		pre, pre1 = bpi, pre
	}
	return pre
}
