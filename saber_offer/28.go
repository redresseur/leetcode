package main

/*
剑指 Offer 28. 对称的二叉树
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/
/*
	标签：二叉树 BFS DFS 对称的二叉树 简单
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 利用 BFS
func isSymmetricByBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	level, stack, next_stack := 0, map[int]*TreeNode{0: root}, map[int]*TreeNode{}
	for len(stack) != 0 {
		// 遍历下一层
		factor := 1 << level
		for i, node := range stack {
			if mirror, ok := stack[factor-i-1]; !ok {
				return false
			} else if node.Val != mirror.Val {
				return false
			}

			if node.Left != nil {
				next_stack[2*i] = node.Left
			}

			if node.Right != nil {
				next_stack[2*i+1] = node.Right
			}
		}
		stack, next_stack = next_stack, map[int]*TreeNode{}
		level++
	}

	return true
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func isSymmetricByDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root, root)
}

func isSymmetricByDFSandStack(root *TreeNode) bool {
	if root == nil {
		return true
	}

	q := []*TreeNode{root, root}
	for len(q) != 0 {
		l, r := q[0], q[1]
		q = q[2:] // 弹出栈顶
		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		q = append(q, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}

/*
	使用栈模拟递归：
	一般流程：
	初始化栈：stack{root}
	for stack.len() != 0:
		第一步：出栈
		node := stack.pop()
		第二步：执行操作
		operator(node) // 此处可能退出 或 跳过循环
		第三步：下一循环 元素 入栈
		通常为：stack.push(node.Right, node.Left) // 注意入栈顺序，后进先出

	针对树的 BFS 属于先序遍历：
	规则：扫描过的则不再扫描
	扫描秒本节点:
		沿着左子树前进:
			遇到叶子节点，next = 兄弟节点
		如果兄弟节点的子树扫描完成:
			返回父节点，next = 父节点->兄弟节点

	扫描过的则不再扫描：不在栈中存在则不会再被扫描
	沿着左子树前进：将左子树放在栈顶
	叶子节点，下一个为兄弟节点：将右子树在左子树前压入栈
		      1
			/  \
			2   2
			\   \
			3    3
	[1]
	[2,2]
	[2,3,null]
	[2,3]
	[2]
	[3,null]
	[3]
	[]
*/
