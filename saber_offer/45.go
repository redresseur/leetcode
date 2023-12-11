package main

import (
	"math"
	"sort"
	"strconv"
)

/*
解题思路
传递性证明：

xy < yx, yz < zy => xc < cx

x 十进制 有 a 位， y 十进制有 b 位, z 十进制有 c 位

x*10^b + y < y*10^a + x
x/(10^a - 1) < y/(10^b - 1)
y/(10^b - 1) < c/(10^c - 1)

x/(10^a - 1) < c/(10^c - 1)
xc < cx

作者：luckydogchina
链接：https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/by-luckydogchina-8ynm/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func dec(i int) float64 {
	ii, n := i, 0
	for {
		n++
		if ii = ii / 10; ii == 0 {
			break
		}
	}
	return float64(i) / (math.Pow10(n) - 1)
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		return dec(x) < dec(y)
	})

	res := ""
	for _, s := range nums {
		res += strconv.Itoa(s)
	}
	return res
}
