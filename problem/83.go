package main

/*
83. 删除排序链表中的重复元素 tag: 链表
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 核心思想，仍然是前哨指针
	prev, next := head, head.Next
	for next != nil {
		if prev.Val == next.Val {
			prev.Next = next.Next
		} else {
			prev = next
		}

		next = next.Next
	}

	return head
}
