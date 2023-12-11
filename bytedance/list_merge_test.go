package bytedance

import (
	"fmt"
	"strconv"
	"testing"
)

func array2List(arr []int) *listNode {
	if len(arr) == 0 {
		return nil
	}

	head := &listNode{}
	for next, i := head, 0; i < len(arr); i++ {
		next.Val = arr[i]
		if i == len(arr)-1 {
			break
		}
		next.Next = &listNode{}
		next = next.Next
	}
	return head
}

func sprintList(list *listNode) string {
	str := "["
	for list != nil {
		str += strconv.Itoa(list.Val) + ","
		list = list.Next
	}
	str += "]"
	return str
}

func TestMerge(t *testing.T) {
	A, B := merge(array2List([]int{3, 5, 10, 25}), array2List([]int{1, 5, 5, 10, 35}))
	fmt.Println(sprintList(A), sprintList(B))
	A, B = merge(array2List([]int{}), array2List([]int{}))
	fmt.Println(sprintList(A), sprintList(B))
	A, B = merge(array2List([]int{1, 1, 2, 5}), array2List([]int{1, 1, 1, 1}))
	fmt.Println(sprintList(A), sprintList(B))
}
