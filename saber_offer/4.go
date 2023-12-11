package main

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。

二分查找 分治 矩阵

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 核心思路: 模仿排序二叉树（左小，右大），以右上角为根节点，大于 root 则
// 右移（移动到下一行），小于root 则向左移动（移动到左一行）
// 考察二分思想：找到一个分界点和一个指针（一维或者二维）移动方案，
// 每一次判断都可以缩小一般的查找范围
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n, m := len(matrix), 0
	if n > 0 {
		m = len(matrix[0])
	}

	if n < 0 || m < 0 {
		return false
	}

	for line, column := 0, m-1; column >= 0 && line < n; {
		v := matrix[line][column]
		if target == v {
			return true
		} else if target < v {
			column--
		} else if target > v {
			line++
		}
	}

	return false
}
