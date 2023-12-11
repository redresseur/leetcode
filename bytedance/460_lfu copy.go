package bytedance

import (
	"container/list"
)

// func init() { debug.SetGCPercent(-1) }

type lfuObj struct {
	k, v, freq int
}

type LFUCache struct {
	freqs        map[int]*list.List
	indexs       map[int]*list.Element
	cap, minFreq int
}

func ConstructorLFU(capacity int) LFUCache {
	lfu := LFUCache{
		indexs: make(map[int]*list.Element),
		freqs:  make(map[int]*list.List),
		cap:    capacity,
	}
	return lfu
}

func (this *LFUCache) Get(key int) int {
	v := -1
	if obj := this.get(key); obj != nil {
		v = obj.v
	}
	return v
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if obj := this.get(key); obj != nil {
		obj.v = value
		return
	}
	if this.cap == len(this.indexs) { // 驱逐key
		this.remove(this.freqs[this.minFreq].Back())
	}
	this.indexs[key] = this.freqList(1).PushFront(&lfuObj{k: key, v: value, freq: 1}) // 创建索引
	this.minFreq = 1
	return
}

// 移除 元素
func (this *LFUCache) remove(e *list.Element) *lfuObj {
	// 从旧链表中移除
	obj := e.Value.(*lfuObj)
	lfreq := this.freqList(obj.freq)
	if lfreq.Remove(e); this.minFreq == obj.freq && lfreq.Len() == 0 {
		this.minFreq = obj.freq + 1 // 更新最小 freq
	}
	delete(this.indexs, obj.k) // 清理索引
	return obj
}

// 获取 链表
func (this *LFUCache) freqList(freq int) *list.List {
	// 加入新链表
	lfreq := this.freqs[freq]
	if lfreq == nil {
		lfreq = list.New()
		this.freqs[freq] = lfreq
	}
	return lfreq
}

// 获取 元素
func (this *LFUCache) get(key int) *lfuObj {
	evalue := this.indexs[key]
	if evalue == nil {
		return nil
	}
	// 从旧链表中移除
	obj := this.remove(evalue)
	obj.freq++
	this.indexs[obj.k] = this.freqList(obj.freq).PushFront(obj) // 更新索引
	return obj
}
