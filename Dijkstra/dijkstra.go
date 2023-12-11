package dijkstra

import (
	"container/heap"
	"fmt"
)

// 一个有向边
type Edge struct {
	start, end int
	weight     int
}

type Vertex struct {
	id   int // 顶点编号
	dist int // 距离起点最近距离
}

type Graph struct {
	adj map[int][]Edge // 邻接表
}

func (g *Graph) addEdge(start, end, weight int) {
	g.adj[start] = append(g.adj[start], Edge{start: start,
		end: end, weight: weight}) // 添加边
}

type minTop []*Vertex // 存储各个节点到 起点距离的 最小堆

func (mt *minTop) Push(x interface{}) { // add x as element Len()
	*mt = append(*mt, x.(*Vertex))
}

func (mt *minTop) Pop() interface{} { // remove and return element Len() - 1.
	if len(*mt) == 0 {
		return nil
	}

	tail := (*mt)[len(*mt)-1]
	*mt = (*mt)[:len(*mt)-1] // 弹出队列
	return tail
}

func (mt *minTop) Less(i, j int) bool {
	return (*mt)[i].dist < (*mt)[j].dist
}

// Swap swaps the elements with indexes i and j.
func (mt *minTop) Swap(i, j int) {
	(*mt)[i], (*mt)[j] = (*mt)[j], (*mt)[i]
}

func (mt *minTop) Len() int {
	return len(*mt)
}

// 保证每一步扩展的节点，是在所有剩余节点里距离起点最近的
// 时间复杂度：O(E*logV)，最坏可能性是遍历了所有的边，每一个更新 vertex minTop 时间复杂度为 O(logV)。
// 输入：start， end
// 输出：shortest
func (g *Graph) dijkstra(start, end int) int {
	shortest := -1 // 初始化为 -1

	queue := minTop{}         // 记录各个点到 start 的距离集合
	inqueue := map[int]bool{} // 记录节点是否在 queue 中，减少操作时间

	getMin := func() int { // 获取距离起点最近的那个点
		return 0
	}

	getVertex := func(id int) int { // 获取指定 id 的 queue 索引
		if inqueue[id] {
			for i := range queue {
				if queue[i].id == id {
					return i
				}
			}
		}
		return -1
	}

	addVertex := func(id, dist int) { // 将节点 id 和 dist 加入 queue
		inqueue[id] = true
		heap.Push(&queue, &Vertex{id: id, dist: dist})
		return
	}

	updateVertex := func(index int, dist int) { // 更新 queue[index] 的 dist
		queue[index].dist = dist
		heap.Fix(&queue, index)
		return
	}

	removeVertex := func(index int) { // 移除指定的 queue[index]
		delete(inqueue, queue[index].id)
		heap.Remove(&queue, index)
		return
	}

	sizeVertex := queue.Len // 获取 queue 的长度

	predecessor := map[int]int{} // 记录最优路径，<目的 id，上一步 id>
	var printPath func(end int)  // 输出最优路径
	printPath = func(end int) {
		if end == start {
			fmt.Print(end)
			return
		}
		if prev, ok := predecessor[end]; ok {
			printPath(prev)
			fmt.Print("->", end)
		}
	}

	addVertex(start, 0) // 初始化 queue，加入起点
	for sizeVertex() != 0 {
		min := queue[getMin()]
		if min.id == end {
			shortest = min.dist
			break
		}
		for _, e := range g.adj[min.id] {
			dist := min.dist + e.weight
			if vi := getVertex(e.end); vi < 0 {
				addVertex(e.end, dist)      // 新建距离记录
				predecessor[e.end] = min.id // 更新 e.end 最小 dist 的上级节点
			} else if dist < queue[vi].dist { // 与原有 dist 进行比较
				updateVertex(vi, dist)      // 更新距离记录
				predecessor[e.end] = min.id // 更新 e.end 最小 dist 的上级节点
			}
		}
		removeVertex(getVertex(min.id)) // 移除 min，因为它的下一步节点都被遍历了。
	}

	printPath(end)
	return shortest
}
