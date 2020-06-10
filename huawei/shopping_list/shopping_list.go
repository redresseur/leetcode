package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*暴力算法*/
type Goods struct {
	V       int    `desc:"商品价格"`
	W       int    `desc:"商品权重"`
	Up      *Goods `desc:"依赖的主键, nil表示商品为主键"`
	UpIndex int
}

type node struct {
	r map[int]bool
	t int
	v int
}

func newNode(t int) *node {
	return &node{
		r: map[int]bool{},
		t: t,
	}
}

func copyNode(n *node) *node {
	res := &node{
		t: n.t,
		r: map[int]bool{},
	}

	for k, v := range n.r {
		res.r[k] = v
	}

	return res
}

// 动态规划
func ShoppingList1(N int, goods []*Goods) (res int) {
	var dp = []*node{}
	for i, g := range goods {
		gt := g.V * g.W
		for j, lent := 0, len(dp); j < lent; j++ {
			node := dp[j]
			if g.UpIndex >= 0 {
				if !node.r[g.UpIndex] {
					continue
				}
			}
			nk, nt := g.V+node.v, gt+node.t
			if nk > N {
				continue
			}

			nN := copyNode(node)
			nN.r[i] = true
			nN.t = nt
			nN.v = nk
			dp = append(dp, nN)
		}

		if g.UpIndex >= 0 {
			continue
		}

		nN := newNode(gt)
		nN.r[i] = true
		nN.v = g.V
		dp = append(dp, nN)
	}

	for _, v := range dp {
		if v.t > res {
			res = v.t
		}
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ns, _, err := reader.ReadLine()
	if err != nil {
		return
	}
	arrs := strings.Split(string(ns), " ")
	if len(arrs) < 2 {
		return
	}

	N, err := strconv.Atoi(arrs[0])
	if err != nil {
		panic(err)
	}

	m, err := strconv.Atoi(arrs[1])
	if err != nil {
		panic(err)
	}

	gs := []*Goods{}
	for i := 0; i < m; i++ {
		ns, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}

		arrs := strings.Split(string(ns), " ")
		if len(arrs) < 3 {
			panic(err)
		}

		v, err := strconv.Atoi(arrs[0])
		if err != nil {
			panic(err)
		}

		w, err := strconv.Atoi(arrs[1])
		if err != nil {
			panic(err)
		}

		up, err := strconv.Atoi(arrs[2])
		if err != nil {
			panic(err)
		}
		gs = append(gs, &Goods{V: v, W: w, UpIndex: up - 1})
	}

	fmt.Println(ShoppingList1(N, gs))
}
