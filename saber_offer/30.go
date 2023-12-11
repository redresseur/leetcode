package main

type MinStack struct {
	stack, minStack []int
}

/** initialize your data structure here. */
func Constructor30() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	// 添加最小值
	if len(this.minStack) == 0 || x < this.stack[this.minStack[len(this.minStack)-1]] {
		this.minStack = append(this.minStack, len(this.stack)-1)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	// 检查最小值记录是否需要变动
	if len(this.stack)-1 <= this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	// 弹出栈
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.minStack) <= 0 {
		return 0
	}
	return this.stack[this.minStack[len(this.minStack)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
