package bytedance

/*
143. 重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
https://leetcode.cn/problems/reorder-list/
*/

/**
解题思路很简单，但是有很多细节需要挖掘：

反转链表的时候需要在头节点加一个哨兵
快慢指针的循环退出条件
合并后的新链 tail.Next 一定要赋值 NULL, 否则会形成环
下面讲一下快慢指针找链表中点的循环退出条件，第一种：


    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

奇数链表：
    1->2->3->4->5->NULL
          s         f

偶数链表：
    1->2->3->4->NULL
          s      f
第二种：


    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
奇数链表：
    1->2->3->4->5
          s     f
偶数链表：
    1->2->3->4->NULL
       s  f
总结一般规律：
奇数链表：无论使用哪种循环，最后 slow 指向中间节点，fast 指向 tail 节点，因为其退出循环都是因为 fast.Next != nil ;
偶数链表：

第一种循环，slow 指向中间偏右，fast 指向 NULL，退出循环条件为 fast!=nil
第二种循环，slow 指向中间偏左，fast 指向 tail 前一个节点，退出循环的条件为 fast.Next.Next!=nil
而本题中，我们找到新链表的尾结点，并且要赋值 new_tail.Next==NULL。综上选择了第二种循环：
1. 奇数链表，则slow 指针最后指向的即为 new_tail；
2. 偶数链表，则合并后指向 new_tail
同时从原始链表的 slow.Next 开始反转链表。

作者：luckydogchina
链接：https://leetcode.cn/problems/reorder-list/solution/fan-by-luckydogchina-e9b3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func reverseList(head *ListNode) *ListNode {
	start := &ListNode{}
	rh := start
	for head != nil {
		next, rhNext := head.Next, rh.Next
		rh.Next, head.Next = head, rhNext
		head = next
	}
	return start.Next
}

func reorderList(head *ListNode) {
	slow, fast := head, head
	// 快慢指针
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	nhead, rl := head, reverseList(slow.Next)
	slow.Next = nil
	for rl != nil { // 合并链表
		next, rlNext := nhead.Next, rl.Next
		nhead.Next, rl.Next = rl, next
		nhead, rl = next, rlNext
	}
	return
}
