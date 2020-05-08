package binary_search

/*
	搜索旋转排序数组
*/

// 查找循环有序数组中的最小值索引
// [4,5,6,1,2,3]则最小值索引为 3
func searchLittle(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	} else if numsLen == 1 {
		return 0
	}

	low, high, mid := 0, numsLen-1, 0
	for low <= high {
		mid = low + (high-low)>>1
		if mid != 0 && nums[mid-1] > nums[mid] {
			return mid
		} else {
			if nums[mid] < nums[0] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}

	return 0
}

// 解題思路，先找到跳變位置，然后再搜索
func SearchIndex(nums []int, target int) int {
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

	low, high, mid := 0, numsLen-1, -1
	if index := searchLittle(nums); index != 0 {
		if target >= nums[0] {
			high = index - 1
		} else {
			low = index
		}
	}

	for low <= high {
		mid = low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 解题思索, 根据跳变特性
