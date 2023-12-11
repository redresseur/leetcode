package bytedance

// // import (
// // 	"container/list"
// // )

// type lfuObj struct {
// 	k, v, c int
// }

// type LFUCache struct {
// 	counters map[int]*list.Element
// 	indexs   map[int]*list.Element
// 	lnodes   *list.List
// 	cap      int
// }

// func ConstructorLFU(capacity int) LFUCache {
// 	lfu := LFUCache{
// 		lnodes:   list.New(),
// 		indexs:   make(map[int]*list.Element),
// 		counters: make(map[int]*list.Element),
// 		cap:      capacity,
// 	}
// 	return lfu
// }

// func (this *LFUCache) get(key int) *list.Element {
// 	return this.indexs[key]
// }

// func (this *LFUCache) update(e *list.Element, value int) {
// 	e.Value.(*lfuObj).v = value
// 	this.updateEle(e)
// }

// func (this *LFUCache) updateEle(e *list.Element) {
// 	obj := e.Value.(*lfuObj)
// 	// 更新 counters[obj.c]
// 	if this.counters[obj.c] == e {
// 		// prev 可能为下一个 counter 或者为 nil
// 		if prev := e.Prev(); prev == nil ||
// 			prev.Value.(*lfuObj).c != obj.c {
// 			delete(this.counters, obj.c)
// 		} else {
// 			this.counters[obj.c] = prev
// 		}
// 	}
// 	obj.c++
// 	// 更新 counters[obj.c+1]
// 	mark := this.counters[obj.c]
// 	if mark == nil {
// 		mark = this.counters[obj.c-1]
// 	}
// 	if mark != nil {
// 		this.lnodes.MoveAfter(e, mark)
// 	}
// 	this.counters[obj.c] = e
// }

// func (this *LFUCache) Get(key int) int {
// 	evalue := this.get(key)
// 	if evalue == nil {
// 		return -1
// 	}
// 	this.updateEle(evalue)
// 	return evalue.Value.(*lfuObj).v
// }

// func (this *LFUCache) evict() {
// 	head := this.lnodes.Front()
// 	obj := head.Value.(*lfuObj)
// 	// 删除访问索引
// 	delete(this.indexs, obj.k)
// 	// 更新对应 counters[c], 如果索引对应的节点为 head
// 	// 说明该 counters[c] 只存在一个 节点
// 	if this.counters[obj.c] == head {
// 		delete(this.counters, obj.c)
// 	}
// 	// 删除 head
// 	this.lnodes.Remove(head)
// }

// func (this *LFUCache) insert(key, value int) {
// 	if len(this.indexs) == this.cap {
// 		this.evict()
// 	}
// 	obj := lfuObj{k: key, v: value, c: 1}
// 	// 更新 counters[1] 的索引
// 	e := this.counters[1]
// 	if e == nil {
// 		e = this.lnodes.PushFront(&obj)
// 	} else {
// 		e = this.lnodes.InsertAfter(&obj, e)
// 	}
// 	this.indexs[key], this.counters[1] = e, e
// }

// func (this *LFUCache) Put(key int, value int) {
// 	evalue := this.get(key)
// 	if evalue != nil {
// 		this.update(evalue, value)
// 	} else {
// 		this.insert(key, value)
// 	}
// 	return
// }

// /**
//  * Your LFUCache object will be instantiated and called as such:
//  * obj := Constructor(capacity);
//  * param_1 := obj.Get(key);
//  * obj.Put(key,value);
//  */
