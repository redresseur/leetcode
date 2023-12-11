package main

type CQueue struct {
	in, out     []int
	in_c, out_c int
}

func Constructor() CQueue {
	return CQueue{
		in:  make([]int, 10001),
		out: make([]int, 10001),
	}
}

func (cq *CQueue) AppendTail(value int) {
	cq.in[cq.in_c] = value
	cq.in_c++
}

func (cq *CQueue) DeleteHead() int {
	head := -1
	if cq.out_c > 0 {
		head = cq.out[cq.out_c-1]
		cq.out_c--
		return head
	}

	if cq.in_c <= 0 {
		return head
	}

	for cq.in_c > 0 {
		cq.out[cq.out_c] = cq.in[cq.in_c-1]
		cq.out_c++
		cq.in_c--
	}

	head = cq.out[cq.out_c-1]
	cq.out_c--
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
