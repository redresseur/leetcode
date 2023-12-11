package main

/*
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// tag: 数学 双指针 字符串
func reverseLeftWords(s string, n int) string {
	s_len := len(s)
	if n <= 0 || s_len <= n {
		return s
	}

	rs := []rune(s)
	for old, tmp, count := 0, 0, 0; ; {
		new := (s_len + old - n) % s_len
		if old < n {
			count++
		}

		/*
			a   b | c d e f
			[e] b | c d a f
			c   b | e d a f
			c [f] | e d a b
			c  d  | e f a b
			核心思想：使用 old 作为临时寄存器，
		*/
		if tmp == new {
			if count == n {
				break
			}

			tmp++
			old = tmp
		} else {
			rs[tmp], rs[new] = rs[new], rs[tmp]
			old = new
		}
	}

	return string(rs)
}
