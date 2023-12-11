package main

/*
剑指 Offer 25. 合并两个排序的链表 tag: 链表
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

示例1：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
限制：

0 <= 链表长度 <= 1000

注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
/*
解题思路
添加一个前哨节点，用于操作其 next 指针，也方便返回结果。
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{}
	merge := start
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1, start.Next = l1.Next, l1
		} else {
			l2, start.Next = l2.Next, l2
		}
		start = start.Next
	}

	if l1 != nil {
		start.Next = l1
	} else {
		start.Next = l2
	}

	return merge.Next
}
