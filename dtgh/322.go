package dtgh

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。
示例 1:
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:
输入: coins = [2], amount = 3
输出: -1

说明:
你可以认为每种硬币的数量是无限的
*/
// 状态迁移方程：F(a) = min(F(a - coin1),a - coin2),... F(a - coinN))
func coinChange(coins []int, amount int) int {
	bp := make([]int, amount+1, amount+2)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if j := i - coin; j > 0 && bp[j] > 0 {
				if bp[i] > bp[j]+1 || bp[i] == 0 {
					bp[i] = bp[j] + 1
				}
			} else if j == 0 {
				bp[i] = 1
			}
		}
	}
	if amount > 0 && bp[amount] == 0 {
		return -1
	}
	return bp[amount]
}
