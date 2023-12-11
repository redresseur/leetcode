package bytedance

import "math"

/*
使用 BFS 的方式去遍历每一层树，用 map 记录遍历过的情况进行去重剪枝：


	   8        queue: 8
	 /   \
    7     4     queue: 7, 4
   / \   /  \
  6   3  3   0  queue: 6, 3 ,0
当访问到 0 节点时则停止遍历，此处使用的是 visted[0] == true 来判断。

时间复杂度：O(n*sqrt(n))。
空间复杂度：O(n)

作者：luckydogchina
链接：https://leetcode.cn/problems/perfect-squares/solution/by-luckydogchina-9sx8/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func numSquaresByDFS(n int) (min int) {
	queue, next_queue, visted := []int{n}, []int{}, make([]bool, n+1)
	for min = 0; len(queue) != 0 && !visted[0]; min++ {
		for i := 0; i < len(queue) && !visted[0]; i++ {
			for j := 1; j*j <= queue[i]; j++ {
				if diff := queue[i] - j*j; !visted[diff] {
					visted[diff], next_queue = true, append(next_queue, diff)
				}
			}
		}
		queue, next_queue = next_queue, queue[:0]
	}

	return min
}

func numSquaresByDp(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			if min > dp[i-j*j] {
				min = dp[i-j*j]
			}
		}
		dp[i] = min + 1
	}

	return dp[n]
}
