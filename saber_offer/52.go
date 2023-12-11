package main

/*
	剑指 Offer 52. 两个链表的第一个公共节点 tag: 链表 双指针
	核心思路：「普通人版本」对齐两个链表，长的链表先走出 l_long - l_short 步
*/
func lenList(head *ListNode) (l int) {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		l += 2
	}
	if fast != nil {
		l += 1
	}
	return l
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 对齐长度
	lenA, lenB := lenList(headA), lenList(headB)
	if k := lenA - lenB; k > 0 {
		for headA != nil && k > 0 {
			headA = headA.Next
			k--
		}
	} else if k < 0 {
		for headB != nil && k < 0 {
			headB = headB.Next
			k++
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA, headB = headA.Next, headB.Next
	}

	return nil
}

/*
	「大佬版本」 假设 链表 A 长度为 a + c , B 长度 为 b + c，c 为公共长度
	当 p1 和 p2 指针都 走 a + b + c 长度后会「浪漫」相遇
	p1:
    1 - 2 -3 - 4\
			     - 5 - 6 - 7
		   8 - 9/
		   p2:

    1 - 2 -3 - 4\  p1:
			     - 5 - 6 - 7
		   8 - 9/         p2:

   p2:                        p2 指针走到尽头后，会指向 A 的头部
	1 - 2 -3 - 4\      p1:
			     - 5 - 6 - 7
		   8 - 9/

		p2:
	1 - 2 - 3 - 4\         p1:
			      - 5 - 6 - 7
			8 - 9/

	        p2:             p1 指针走到尽头后，会指向 B 的头部
	1 - 2 - 3 - 4\
			      - 5 - 6 - 7
			8 - 9/
		   p1:

	1 - 2 - 3 - 4\ p2:      在 p1 和 p2 都走过 a + b + c 步后会相遇
			      - 5 - 6 - 7
			8 - 9/ p1:

	Note: 如果两个链条没有公共部分，则最终 p1 和 p2 在 走过 a+b 步后会同时指向 nil
	super romantic ！！！！
*/
func getIntersectionNodeOpz(headA, headB *ListNode) *ListNode {
	// 对齐长度
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}

		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}

	return p1
}
