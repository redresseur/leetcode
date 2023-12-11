package main

/*
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例 1:

输入：s = "abaccdeff"
输出：'b'
示例 2:

输入：s = ""
输出：' '


限制：

0 <= s 的长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
	核心思想：因为字符串中只包含 a-z 26个小写字符，所以可以使用 [26]int32{} 数组记录状态，
	使用 int32 是因为 int16 无法容纳 50000. 每一个位置有三个状态 0 没有出现 -1 重复出现
	>0 记录对应字符第一次出现的 index+1.
	时间复杂度：O(n), 空间复杂度：O(1)
*/
func firstUniqChar(s string) byte {
	cacl := [26]int32{}
	for index, r := range []byte(s) {
		if i := r - 'a'; cacl[i] == 0 {
			cacl[i] = int32(index + 1)
		} else if cacl[i] > 0 {
			cacl[i] = -1
		}
	}

	ci, index := -1, int32(1<<15-1)
	for i, si := range cacl {
		if si > 0 && index > si {
			index, ci = si, i
		}
	}

	if ci < 0 {
		return ' '
	}

	return 'a' + byte(ci)
}
