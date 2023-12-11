package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
注意到方法一需要使用哈希表记录每一个节点对应新节点的创建情况，而我们可以使用一个小技巧来省去哈希表的空间。

我们首先将该链表中每一个节点拆分为两个相连的节点，例如对于链表 A \rightarrow B \rightarrow CA→B→C，我们可以将其拆分为 A \rightarrow A' \rightarrow B \rightarrow B' \rightarrow C \rightarrow C'A→A
′
 →B→B
′
 →C→C
′
 。对于任意一个原节点 SS，其拷贝节点 S'S
′
  即为其后继节点。

这样，我们可以直接找到每一个拷贝节点 S'S
′
  的随机指针应当指向的节点，即为其原节点 SS 的随机指针指向的节点 TT 的后继节点 T'T
′
 。需要注意原节点的随机指针可能为空，我们需要特别判断这种情况。

当我们完成了拷贝节点的随机指针的赋值，我们只需要将这个链表按照原节点与拷贝节点的种类进行拆分即可，只需要遍历一次。同样需要注意最后一个拷贝节点的后继节点为空，我们需要特别判断这种情况。

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/solution/fu-za-lian-biao-de-fu-zhi-by-leetcode-so-9ik5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func copyRandomList1(head *Node) (res *Node) {
	if head == nil {
		return nil
	}

	// 复制自身为后继节点
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}

	// 为自身的复制节点复制 random 指针
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	// 将链一分为二
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil { // 未到末尾
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew
}

func copyRandomList(head *Node) (res *Node) {
	iter, num := head, 0
	copyList := make([]Node, 0, 16)
	randIndex := make(map[*Node]int, 16)

	// 复制链表
	for num := 0; iter != nil; num++ {
		copyList = append(copyList, *iter)
		randIndex[iter] = num
		iter = iter.Next
	}

	// 恢复指针
	for i := range copyList {
		index, ok := randIndex[copyList[i].Random]
		if ok {
			copyList[i].Random = &copyList[index]
		}
		if i < num-1 {
			copyList[i].Next = &copyList[i+1]
		}
	}

	if num > 0 {
		res = &copyList[0]
	}

	return
}
