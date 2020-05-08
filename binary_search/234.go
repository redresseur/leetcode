package binary_search

/*
	234 回文链表
	https://leetcode-cn.com/problems/palindrome-linked-list/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome1(head *ListNode) (res bool) {
	if head == nil || head.Next == nil {
		return true
	}

	// 寻找中间节点, 反转前半段
	reverse, tmp := (*ListNode)(nil), (*ListNode)(nil)
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		tmp = slow
		slow = slow.Next
		fast = fast.Next.Next

		tmp.Next = reverse
		reverse = tmp
	}

	// 判断长度奇偶，非空为奇数，空为偶数
	if nil != fast {
		slow = slow.Next
	}

	// 比较
	for ; slow != nil && reverse != nil; slow, reverse = slow.Next, reverse.Next {
		if slow.Val != reverse.Val {
			return false
		}
	}

	return true
}

func isPalindrome(head *ListNode) (res bool) {
	if head == nil || head.Next == nil {
		return true
	}

	reverse, tmp := (*ListNode)(nil), (*ListNode)(nil)
	fast, slow := head, head

	// 寻找中间节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 判断长度奇偶，非空为奇数，空为偶数
	if nil != fast {
		slow = slow.Next
	}

	// 反转后半段
	for slow != nil {
		tmp = slow
		slow = slow.Next
		tmp.Next = reverse
		reverse = tmp
	}

	// 比较
	for ; head != nil && reverse != nil; head, reverse = head.Next, reverse.Next {
		if head.Val != reverse.Val {
			return false
		}
	}

	return true
}
