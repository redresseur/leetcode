package bytedance

/*
46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

https://leetcode.cn/problems/permutations/

最基本的思路就是使用递归暴力枚举所有情况，但是要保证每个节点不能重复出现，所以用了一个 map[int]bool 去完成剪枝，为了保证加入的顺序，使用 []int 模拟栈记录。

核心的两个操作 入栈 和 出栈：

	+ 入栈：插入 map 和 压入 arr
	+ 出栈：从 arr 尾部弹出 和 从 map 中移除
官方的思路更加巧妙，通过 first 索引来将数组隔离为 组合|未组合 两部分，然后再针对 未组合 部分去重复操作。本解法的思路更"笨拙"一些，但是胜在更容易理解:

+ 时间复杂度：O(N!)
+ 空间复杂度：O(N)

作者：luckydogchina
链接：https://leetcode.cn/problems/permutations/solution/di-gui-jianzhi-by-luckydogchina-jp8f/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	dp, n := [][]int{}, len(nums)
	path, one := map[int]bool{}, []int{}
	push := func(n int) {
		path[n], one = true, append(one, n)
	}
	pop := func() {
		delete(path, one[len(one)-1])
		one = one[:len(one)-1]
	}
	dump := func() []int {
		onec := make([]int, len(one))
		copy(onec, one)
		return onec
	}

	var backtrack func()
	backtrack = func() {
		if len(one) == n {
			dp = append(dp, dump())
			return
		}

		for i := 0; i < n; i++ {
			if path[nums[i]] {
				continue
			}
			push(nums[i])
			backtrack()
			pop()
		}
	}
	backtrack()
	return dp
}
