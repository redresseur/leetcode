package bytedance

import "math"

/*
7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

https://leetcode.cn/problems/reverse-integer/
*/
func reverse(x int) (res int) {
	maxdiv, mindiv := math.MaxInt32/10, math.MinInt32/10
	for x != 0 {
		if res > maxdiv || res < mindiv {
			return 0
		}

		res *= 10
		next := x % 10
		if next < 0 && next < math.MinInt32-res {
			return 0
		}

		if next > 0 && next > math.MaxInt32-res {
			return 0
		}

		res += next
		x /= 10
	}
	return
}
