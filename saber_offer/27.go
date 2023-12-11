package main

/*
剑指 Offer 27. 二叉树的镜像
请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1



示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]


限制：

0 <= 节点个数 <= 1000

注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/

tag: DFS BFS 二叉树 遍历 二叉树的镜像
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归实现：通过递归 DFS 遍历二叉树
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	res := &TreeNode{Val: root.Val}
	res.Right = mirrorTree(root.Left)
	res.Left = mirrorTree(root.Right)

	return res
}

// 栈实现：通过循环实现
func mirrorTreeByStack(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		// 出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}

		// 入栈, 先右后左
		stack = append(stack, node.Right, node.Left)
		node.Right, node.Left = node.Left, node.Right
	}

	return root
}
