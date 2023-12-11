package main

/*
剑指 Offer 26. 树的子结构
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

     3
    / \
   4   5
  / \
 1   2
给定的树 B：

   4
  /
 1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

输入：A = [1,2,3], B = [3,1]
输出：false
示例 2：

输入：A = [3,4,5,1,2], B = [4,1]
输出：true
限制：

0 <= 节点个数 <= 10000
*/
// 标签：DFS 二叉树 树的子结构
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	核心思路：和在二叉树中遍历 目标值一样的原理：
	func: IsSubStructure(rootA, rootB):
		if checkSubStructure(rootA, rootB):
			return true
		if IsSubStructure(rootA->left, rootB)
			return ture
		if IsSubStructure(rootA->right, rootB)
			return ture
		return false
	这里面我们可以把 checkSubStructure 理解成 "==" 操作，这样就好理解了。

	总结心得：本题与 27 题（二叉树镜像翻转），相同点都是针对最外层的（即定义）函数去递归
	一般而言 28 题也应该如此，但是 28 题需要拆分成两个指针，同时滚动两个指针，而定义函数
	入口只有一个参数，则只好去递归子函数。

	27，26（本）题都是使用一个指针去扫描整个树：
		递归传入参数是一个子树或整树 "root"，另一个参数是目标值，或无 其他参数;
	28题同时滚动两个指针去扫描整个数：
		递归传入参数是两个子树或整树 "root";
*/
func checkSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val { // 查找子结构，不是子树
		return false
	}

	return checkSubStructure(A.Left, B.Left) &&
		checkSubStructure(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	return checkSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
