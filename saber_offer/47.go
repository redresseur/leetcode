package main

/*
剑指 Offer 47. 礼物的最大价值
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？



示例 1:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物


提示：

0 < grid.length <= 200
0 < grid[0].length <= 200
*/
// 标签： 动态规划 剑指 Offer 47. 礼物的最大价值
/*
	状态转移方程：max(i,j) = grid[i][j] + max(max(i,j-1), max(i-1)(j))
	注解：有些算法喜欢用grid本身记录中间状态，来节省空间，但本人不建议这做，一旦写
	不好会很难找bug。
*/
func maxValue(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			u := dp[j]
			if j-1 >= 0 && u < dp[j-1] {
				u = dp[j-1]
			}
			dp[j] = u + grid[i][j]
		}
	}

	return dp[len(dp)-1]
}
