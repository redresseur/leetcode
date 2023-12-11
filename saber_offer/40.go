package main

func getLeastNumbersQuickSelect(arr []int, k int) []int {
	if k <= 0 || len(arr) <= 0 {
		return []int{}
	} else if k >= len(arr) {
		return arr
	}

	c, pivot := arr[len(arr)-1], len(arr)-1
	var j, i int
	for j, i = 0, 0; i < len(arr)-1; i++ {
		if arr[i] <= c {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}

	arr[pivot], arr[j] = arr[j], arr[pivot]
	pivot = j

	if pivot < k {
		getLeastNumbersQuickSelect(arr[pivot:], k-pivot)
	} else if pivot > k {
		getLeastNumbersQuickSelect(arr[:pivot], k)
	}

	return arr[:k]
}

/*
解题思路
使用堆排序，我建议各位了解这种方法，因为在mysql中 orderby limit 就是采用堆排序进行的优化。
假设 limit k，则维持一个 k 大小的堆，筛选出 最小或最大 k 个元素，然后依次弹出堆。

+ 时间复杂度为 O(Nlog(k));
+ 空间复杂度为 O(k);
整个过程分为两个部分：

+ 建堆：在数组最后插入数组，然后自下向上堆化 O(k)
+ 更新堆：如果新节点小于堆顶节点，用新节点替换堆顶节点，然后自上向下堆化 O(N log(k))

作者：luckydogchina
链接：https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/solution/by-luckydogchina-2u2i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func getLeastNumbersHeap(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	} else if k >= len(arr) {
		return arr
	}

	heap := make([]int, 1, k+1)
	for _, in := range arr {
		heap = insertHeap(heap, in, k)
	}
	return heap[1:]
}

func insertHeap(heap []int, in, k int) []int {
	if len(heap)-1 == k {
		if heap[1] <= in {
			return heap
		}
		// update heap
		heap[1] = in
		for pos, i := 1, 0; pos != i; {
			i = pos
			if 2*i < len(heap) && heap[pos] < heap[2*i] {
				pos = 2 * i
			}

			if 2*i+1 < len(heap) && heap[pos] < heap[2*i+1] {
				pos = 2*i + 1
			}

			heap[pos], heap[i] = heap[i], heap[pos]
		}
	} else {
		// build heap
		heap = append(heap, in)
		for i := len(heap) - 1; i > 1; i /= 2 {
			if heap[i] <= heap[i/2] {
				break
			}
			heap[i], heap[i/2] = heap[i/2], heap[i]
		}
	}
	return heap
}
