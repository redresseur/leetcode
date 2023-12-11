package main

/*
剑指 Offer 34. 二叉树中和为某一值的路径
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/
*/

func pathSum(root *TreeNode, target int) [][]int {
	res, path := [][]int{}, []int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}

		target, path = target-root.Val, append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			if target == 0 { // 符合条件
				cpath := make([]int, len(path))
				copy(cpath, path)
				res = append(res, cpath)
			}
		} else {
			dfs(root.Left, target)
			dfs(root.Right, target)
		}
		path = path[:len(path)-1] // 弹出栈
	}

	dfs(root, target)
	return res
}
