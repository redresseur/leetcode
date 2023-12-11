package main

/*
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。



例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [20,9],
  [15,7]
]

tag: BFS DFS 二叉树 遍历

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return nil
	}

	queue, next_queue := []*TreeNode{root}, []*TreeNode{}
	l2r := true
	for q_l := len(queue); q_l != 0; q_l = len(queue) {
		line := make([]int, q_l)
		for i, node := range queue {
			if l2r {
				line[i] = node.Val
			} else {
				line[q_l-i-1] = node.Val
			}

			if left := node.Left; left != nil {
				next_queue = append(next_queue, left)
			}
			if right := node.Right; right != nil {
				next_queue = append(next_queue, right)
			}
		}
		res = append(res, line)
		queue, next_queue = next_queue, queue
		next_queue, l2r = next_queue[:0], !l2r
	}
	return
}
