package binary_tree

type Node struct {
	right, left *Node
	data        int
	cnt         uint16 // 重複計數
}

func NewNode(data int) *Node {
	node := &Node{data: data, cnt: 1}
	return node
}

type Tree struct {
	root *Node
}

// 先序遍历
func (t *Tree) DLR() {
	t.dlr(t.root)
}

func (t *Tree) dlr(node *Node) {
	if node != nil {
		t.dlr(node.left)
		print("[", node.data, ":", node.cnt, "]")
		t.dlr(node.right)
	}
}

// 返回存在的數量
func (t *Tree) Find(data int) int {
	p := t.root
	for p != nil {
		if p.data > data {
			p = p.left
		} else if p.data < data {
			p = p.right
		} else {
			return int(p.cnt)
		}
	}

	return -1
}

// 插入樹
func (t *Tree) Insert(data int) {
	p := &t.root
	for {
		if *p == nil {
			*p = NewNode(data)
			break
		}

		if (*p).data > data {
			p = &((*p).left)
		} else if (*p).data < data {
			p = &((*p).right)
		} else {
			(*p).cnt++
			break
		}
	}

	return
}

// 树的高度
func (t *Tree) BFSHigh() (high int) {
	if t.root == nil {
		return -1
	}

	layer_list := []*Node{t.root}
	for len(layer_list) != 0 {
		new_layer := []*Node{}
		for _, node := range layer_list {
			if node.right != nil {
				new_layer = append(new_layer, node.right)
			}
			if node.left != nil {
				new_layer = append(new_layer, node.left)
			}
		}
		high++
		layer_list = new_layer
	}

	return high - 1
}

func (t *Tree) DFSHigh() (high int) {
	if t.root == nil {
		return 0
	}
	return t.dfsHigh(t.root) - 1
}

func (t *Tree) dfsHigh(node *Node) int {
	if node == nil {
		return 0
	}

	right, left := t.dfsHigh(node.right), t.dfsHigh(node.left)
	if left > right {
		return left + 1
	}

	return right + 1
}

func (t *Tree) Remove(data int) {
	p, pp := t.root, (**Node)(nil)
	// 查找节点
	for p != nil {
		if p.data > data {
			pp, p = &p.left, p.left
		} else if p.data < data {
			pp, p = &p.right, p.right
		} else {
			break
		}
	}

	if p == nil {
		return
	}

	rp, prp := (*Node)(nil), (*Node)(nil)
	if p.right != nil {
		// 查找右子树最小節點
		for rp = p.right; rp.left != nil; {
			prp, rp = rp, rp.left
		}

		// 移除最小节点
		if prp != nil {
			prp.left, rp.right = nil, p.right
		}
		rp.left = p.left
	} else {
		rp = p.left
	}

	// 替换
	if pp != nil {
		*pp = rp
	} else {
		t.root = rp
	}

	return
}

// 二叉树深度优先遍历，非递归
// 采用栈数据结构：根节点初始栈; 弹出栈 -> 记录数据 -> 右节点入栈 -> 左节点入栈; 循环直到栈为空
func (t *Tree) DfsErgodic() []int {
	queue := []*Node{t.root}
	datas := []int{}
	for length := 1; length != 0; length = len(queue) {
		node := queue[length-1]
		datas = append(datas, node.data)
		queue = queue[:length-1]
		if node.right != nil {
			queue = append(queue, node.right)
		}

		if node.left != nil {
			queue = append(queue, node.left)
		}
	}

	return datas
}

// 二叉树广度优先遍历
// 采用队列的方式,节点入队列-> 节点出队列，左、右子节点入队列，记录每一层的节点，直到队列为空
func (t *Tree) BfsErgodic() []int {
	queue := []*Node{t.root}
	datas := []int{}
	for len(queue) != 0 {
		tmp := []*Node{}
		for _, node := range queue {
			datas = append(datas, node.data)
			if node.left != nil {
				tmp = append(tmp, node.left)
				//datas = append(datas, node.left.data)
			}

			if node.right != nil {
				tmp = append(tmp, node.right)
				//datas = append(datas, node.right.data)
			}
		}

		queue = tmp
	}

	return datas
}
