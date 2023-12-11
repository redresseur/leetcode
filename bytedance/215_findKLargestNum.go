package bytedance

/*
搜索数组中 K 大的数
作者：luckydogchina
链接：https://leetcode.cn/problems/kth-largest-element-in-an-array/solution/kuai-su-sou-suo-by-luckydogchina-72hk/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// 基于堆排序
func findKthLargestByHeap(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k <= 0 || n < k {
		return -1
	}
	// build heap
	heap := make([]int, k+1)
	for i := 0; i < k; i++ {
		heap[i+1] = nums[i]
		for n := i + 1; n > 1; n /= 2 { // down to up
			if heap[n] >= heap[n/2] {
				break
			}
			heap[n], heap[n/2] = heap[n/2], heap[n]
		}
	}
	// find kth largest
	for i := k; i < n; i++ {
		if nums[i] <= heap[1] {
			continue
		}
		heap[1] = nums[i]
		for n, next := 1, 1; n <= k; n = next { // up to down
			if 2*n <= k && heap[next] > heap[2*n] {
				next = 2 * n
			}
			if 2*n+1 <= k && heap[next] > heap[2*n+1] {
				next = 2*n + 1
			}
			if n == next {
				break
			}
			heap[n], heap[next] = heap[next], heap[n]
		}
	}
	return heap[1]
}

// 基于快速查找
func findKthLargest(nums []int, k int) int {
	if n := len(nums); n == 0 || k <= 0 || n < k {
		return -1
	}

	var quicFind func(l, r int) int
	quicFind = func(l, r int) int {
		c, pivot := nums[r], l
		for i := l; i < r; i++ {
			if nums[i] > c {
				nums[i], nums[pivot] = nums[pivot], nums[i]
				pivot++
			}
		}

		nums[r], nums[pivot] = nums[pivot], nums[r]
		if pivot == k-1 {
			return nums[pivot]
		} else if pivot < k-1 {
			return quicFind(pivot+1, r)
		} else {
			return quicFind(l, pivot-1)
		}
	}

	return quicFind(0, len(nums)-1)
}
