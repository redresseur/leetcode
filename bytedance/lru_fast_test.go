package bytedance

import "testing"

func TestLruFast(t *testing.T) {
	cache := &LRUCacheFast{capacity: 2}
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Put(4, 4)
}
