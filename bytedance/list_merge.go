package bytedance

type listNode struct {
	Val  int
	Next *listNode
}

func merge(A, B *listNode) (*listNode, *listNode) {
	startA, startB := &listNode{Next: A}, &listNode{Next: B}
	pa, pb := startA, startB
	for A != nil && B != nil {
		if val := A.Val; val == B.Val {
			for A != nil && val == A.Val {
				pa.Next = A.Next
				A = A.Next
			}

			for B != nil && val == B.Val {
				pb.Next = B.Next
				B = B.Next
			}
		} else if B.Val < A.Val {
			pb, B = B, B.Next
		} else {
			pa, A = A, A.Next
		}
	}

	return startA.Next, startB.Next
}
