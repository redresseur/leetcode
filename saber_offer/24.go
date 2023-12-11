package main

/*
   解题思路：翻转链表最少需要两个指针变量：
   rh: 保存翻转链表的头部
   next: 作为中转存储，进行链表遍历

   过程：
   1）将  head 从 List 中 pop 出来
   2）将 head push 到 resverseList
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var rh, next *ListNode
	for head != nil {
		next = head.Next
		head.Next, rh = rh, head
		head = next
	}

	return rh
}
