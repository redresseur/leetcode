package bytedance

/*
199. 二叉树的右视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
解题思路
广度优先遍历，每层最右的节点值加入结果中。使用两个队列，避免重复申请内存。

时间复杂度：O(N)
空间复杂度：O(N)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	queue, nextQueue := []*TreeNode{root}, []*TreeNode{}
	for len(queue) != 0 {
		res = append(res, queue[len(queue)-1].Val)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue, nextQueue = nextQueue, queue[:0]
	}
	return
}
