package put_apple

import (
	"algo/huawei/put_apple"
	"testing"
)

func TestCount(t *testing.T) {
	t.Log(main.Count(4,3))
	t.Log(main.Count(7,3))
	t.Log(main.Count(3,3))
	t.Log(main.Count(3,4))
	t.Log(main.Count(0,4))
}