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
