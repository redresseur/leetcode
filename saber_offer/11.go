package main

/*
与 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/ 相同
*/
/*
	搜索旋转数组
	二分查找 分治
	核心思路：以最右值为分界点，[最小值,最右值] <= 最右值 且范围内递增，[最左值，最小值）>= 最右值 且范围内递增
	当遇到相等值时可以考虑丢弃最右值向 左 移动一位，剩余的最右值依然满足条件
	https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-by-leetcode-s/
	实现的关键：对于 numbers[pivot] == numbers[high]的处理
	各人理解，为何分区取 low = pivot + 1, high = pivot,
		因为 numbers[pivot] < numbers[high] 时，其可能是最小值，要保留在搜索范围内
		而  numbers[pivot] > numbers[high] 时，其必不能是最小值，无需保留在搜索范围内

	总结：二分查找的高低指针位移规则非常关键，尤其是当 mid == low 情况发生时，一定要保证可以触发循环退出规则。
		此外，不要发生"漏扫" 将 target 值给遗漏掉造成误判。
	tips：为了避免死 mid == low 循环，可以让 high = mid - 1 和 low = mid + 1，但是要小心数组越界。
*/
func minArray(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/solution/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-by-leetcode-s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
