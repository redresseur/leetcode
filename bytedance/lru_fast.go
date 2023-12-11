package bytedance

type listNode1 struct {
	Flag              bool
	Value, Next, Prev int
}

type LRUCacheFast struct {
	index [10001]listNode1

	len, capacity int
}

func (this *LRUCacheFast) moveToBack(key int) {
	if this.index[key+1].Next == 0 {
		return
	}

	// 删除节点
	this.remove(key)
	this.insertToBack(key, this.index[key+1].Value)
	return
}

func (this *LRUCacheFast) insertToBack(key, value int) {
	key += 1
	// 插入队尾
	prev := this.index[0].Prev
	this.index[key].Prev, this.index[key].Next = prev, 0
	this.index[0].Prev, this.index[prev].Next = key, key
	this.index[key].Flag = true
	this.index[key].Value = value
}

// 删除节点
func (this *LRUCacheFast) remove(key int) {
	key += 1
	this.index[key].Flag = false

	prev, next := this.index[key].Prev, this.index[key].Next
	this.index[prev].Next, this.index[next].Prev = next, prev
}

func (this *LRUCacheFast) Get(key int) int {
	if !this.index[key+1].Flag {
		return -1
	}

	this.moveToBack(key)
	return this.index[key+1].Value
}

func (this *LRUCacheFast) Put(key, value int) {
	if this.Get(key) >= 0 {
		this.index[key+1].Value = value
		return
	}

	if this.len == this.capacity {
		this.remove(this.index[0].Next - 1)
	} else {
		this.len++
	}

	this.insertToBack(key, value)
}
