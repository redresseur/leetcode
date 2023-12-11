package bytedance

import "container/list"

type kv struct {
	Key, Val int
}

type LRUCache struct {
	index    [10001]*list.Element
	list     list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node := this.index[key]; node != nil {
		this.list.MoveToBack(node)
		return node.Value.(*kv).Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.index[key]; node != nil {
		node.Value.(*kv).Val = value
		this.list.MoveToBack(node)
		return
	}

	if this.list.Len() == this.capacity { // 驱逐节点
		head := this.list.Front()
		this.list.Remove(head)
		this.index[head.Value.(*kv).Key] = nil
	}

	this.index[key] = this.list.PushBack(&kv{key, value})
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
