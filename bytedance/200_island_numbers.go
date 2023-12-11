package bytedance

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

思路与官方的 dfs 雷同，简单的将就是图的 dfs 遍历，从左上角开始一直遍历到所有边缘节点。为了不重复遍历，将遍历过的节点标记为 '2'(官方标记为'0'，不过都一样的效果不需要纠结。)。
稍微不同的地方在于，为了节省内存，使用栈来模拟递归调用，顺便说一句，用 golang 写算法一定要善用闭包。

作者：luckydogchina
链接：https://leetcode.cn/problems/number-of-islands/solution/dfs-jie-ti-fa-by-luckydogchina-buin/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

时间复杂度：O(MN)
空间复杂度：O(MN)
*/
func numIslands(grid [][]byte) (num int) {
	island, tag, queue := byte('1'), byte('2'), [][2]int{}
	pop := func() [2]int {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		return top
	}

	push := func(top [2]int) {
		if top[0] >= 0 && top[0] < len(grid) &&
			top[1] >= 0 && top[1] < len(grid[top[0]]) &&
			grid[top[0]][top[1]] == island {
			queue = append(queue, top)
		}
	}

	var search func()
	search = func() {
		for len(queue) != 0 {
			// 弹出栈
			top := pop()
			grid[top[0]][top[1]] = tag // 标记该节点
			// 上节点
			push([2]int{top[0], top[1] - 1})
			// 下节点
			push([2]int{top[0], top[1] + 1})
			// 左节点
			push([2]int{top[0] - 1, top[1]})
			// 右节点
			push([2]int{top[0] + 1, top[1]})
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != island {
				continue
			}

			queue = append(queue, [2]int{i, j})
			search()
			num++
		}
	}

	return
}
