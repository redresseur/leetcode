package main

// 整数除法 https://leetcode.cn/problems/xoh6Oh/
// 参考 leetcode 官网的标准答案，将 a，b 输入都转化为 负数，
// 这样可以表达更大的取值范围
// 1) 因为此处假设 int32 的取值范围，所以先过滤掉溢出输入 2^31 -1 >= a,b >= -2^31
// 2) 处理特殊情况：
//   + b 为 0，返回 exp(2^31 -1)
//   + |b| > |a|，返回 0
//   + |b| == |a|，返回 1
// 3) 模拟二进制除法过程:
//    1000111/101
//    1000111
//  - 1010000   0
//    1000111
//  - 0101000   1
//    0011111
//  - 0010100   1
//    0001011
//  - 0001010   1
//    0000001 余数
// 结果为 01110 = 14 == 71/5
// 将除法转化为位移和减法
func divide(a int, b int) int {
	max := 1<<31 - 1
	negitive, result := 0, max // 1 则结果为负
	offset := 0                // a,b 绝对值的最高 bit 位
	a_32, b_32 := int32(a), int32(b)
	if a != int(a_32) || b != int(b_32) {
		goto end
	}

	// 除 0 异常
	if b_32 == 0 {
		goto end
	}

	// 负数表示的值范围更大
	if a_32 > 0 {
		a_32 = -a_32
		negitive++
	}

	if b_32 > 0 {
		b_32 = -b_32
		negitive++
	}

	// |b| > |a| 则为 0
	if b_32 < a_32 || a_32 == 0 {
		result = 0
		goto end
	} else if a_32 == b_32 {
		result = 1
		goto end
	}

	// 此计算 offset 满足 (a >> offset) < b
	// 因为过滤掉了 |a| <= |b| 的情况，所以 b 至少要 << 1 才会比 a 大；
	// 避免位移溢出，此处判断 n 要 < b_32
	for n := b_32 << 1; n >= a_32 && n < b_32; n <<= 1 {
		offset++
	}

	result = 0
	for begin := offset; begin >= 0; begin-- {
		re := a_32 - (b_32 << begin)
		result <<= 1
		if re <= 0 { // 余数不为 0
			result += 1
			a_32 = re
		}
	}

end:
	if negitive == 1 {
		result = -result
	} else if result > max { // 避免结果溢出
		result = max
	}

	return result
}
