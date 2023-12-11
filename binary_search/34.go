package binary_search

/*
	34. 在排序数组中查找元素的第一个和最后一个位置
*/
func searchMin(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	} else {
		if target == nums[0] {
			return 0
		} else if numsLen == 1 {
			return -1
		}
	}

	low, high, mid := 0, numsLen-1, 0
	for low <= high {
		mid = low + (high-low)>>1
		if nums[mid] == target {
			if mid == 0 || nums[mid] != nums[mid-1] {
				return mid
			} else {
				high = mid - 1
			}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	} else {
		if target == nums[0] && numsLen == 1 {
			return []int{0, 0}
		}
	}

	low, high, mid := 0, numsLen-1, 0
	for low <= high {
		mid = low + (high-low)>>1
		if nums[mid] == target {
			if nums[high] != target {
				high = high - 1
			} else if nums[low] != target {
				low = low + 1
			} else {
				break
			}
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low > high {
		return []int{-1, -1}
	}

	return []int{low, high}
}
