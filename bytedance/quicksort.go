package bytedance

import (
	"math/rand"
	"time"
)

// quickSort 使用快速排序算法，排序整型数组
func quickSort(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	// 使用随机数优化快速排序
	rand.Seed(time.Now().Unix())
	r := rand.Intn(b-a) + a

	c := b - 1
	arr[c], arr[r] = arr[r], arr[c]

	j := a
	for i := a; i < c; i++ {
		if arr[i] < arr[c] {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}
	arr[j], arr[c] = arr[c], arr[j]

	quickSort(arr, a, c)
	quickSort(arr, c+1, b)
}

func quickSort2(arr []int, a, b int) {
	if b-a <= 1 {
		return
	}

	j := a
	c := b - 1
	for i := a; i < c; i++ {
		if arr[i] < arr[c] {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
	}

	arr[j], arr[c] = arr[c], arr[j]
	quickSort(arr, a, c)
	quickSort(arr, c+1, b)
}
