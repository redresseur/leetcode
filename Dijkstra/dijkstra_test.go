package dijkstra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstra(t *testing.T) {
	//示意图 https://static001.geekbang.org/resource/image/e2/a9/e20907173c458fac741e556c947bb9a9.jpg?wh=1142*856
	graph := &Graph{map[int][]Edge{}}
	graph.addEdge(0, 1, 10)
	graph.addEdge(0, 4, 15)
	graph.addEdge(1, 3, 2)
	graph.addEdge(1, 2, 15)
	graph.addEdge(2, 5, 5)
	graph.addEdge(3, 5, 12)
	graph.addEdge(3, 2, 1)
	graph.addEdge(4, 5, 10)

	// 1. 正常测试
	assert.Equal(t, 18, graph.dijkstra(0, 5))
	println()

	assert.Equal(t, 10, graph.dijkstra(0, 1))
	println()

	assert.Equal(t, 15, graph.dijkstra(0, 4))
	println()

	assert.Equal(t, 12, graph.dijkstra(0, 3))
	println()

	assert.Equal(t, 13, graph.dijkstra(0, 2))
	println()

	assert.Equal(t, 8, graph.dijkstra(1, 5))
	println()

	// 2. 异常测试
	assert.Equal(t, 0, graph.dijkstra(0, 0))

	assert.Equal(t, -1, graph.dijkstra(5, 0))

	assert.Equal(t, -1, graph.dijkstra(3, 0))

	assert.Equal(t, -1, graph.dijkstra(6, 0))
}
