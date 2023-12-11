package main

/*
1836. 从未排序的链表中移除重复元素 tag: 链表 哈希表
给定一个链表的第一个节点 head ，找到链表中所有出现多于一次的元素，并删除这些元素所在的节点。

返回删除后的链表。
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
	start := &ListNode{Next: head}
	cc := make(map[int]int, 16)
	for head != nil {
		cc[head.Val]++
		head = head.Next
	}

	// 删除链表的核心思想：使用一个前哨指针
	prev, next := start, start.Next
	for next != nil {
		if cc[next.Val] > 1 {
			prev.Next = next.Next // 删除节点，前哨不更新
		} else {
			prev = next // 更新前哨
		}
		next = next.Next
	}

	return start.Next
}
