package bytedance

import "testing"

func TestLru(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
}
