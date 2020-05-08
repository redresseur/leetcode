package binary_search

/*
	搜索旋转排序数组最小值
*/

// # 153
func findMin(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	} else if numsLen == 1 {
		return nums[0]
	}

	low, mid, high := 0, numsLen-1, 0
	for low <= high {
		mid = low + (high-low)>>1
		if mid != 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		} else {
			if nums[mid] >= nums[0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return nums[0]
}
