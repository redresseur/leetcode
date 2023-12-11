package main

func deleteDuplicatesRemoveAll(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	start := &ListNode{Next: head}
	prev, next := start, head
	for next != nil {
		if next.Next == nil || next.Val != next.Next.Val {
			prev = next // 更新前向节点
			next = next.Next
			continue
		}

		for val := next.Val; next != nil; {
			if val != next.Val {
				break
			} else {
				prev.Next = next.Next // 删除节点
				next = next.Next
			}
		}
	}

	return start.Next
}
