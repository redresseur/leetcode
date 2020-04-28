package positivenumbers

//1.Given an array of ints = [6, 4, -3, 5, -2, -1, 0, 1, -9],
//implement a function in any language to move all positive
//numbers to the left, and move all non-positive numbers to the right.
//Try your best to make its time complexity to O(n), and
//space complexity to O(1).

func Positivenum(ints []int)[]int{
	for low, high := 0, len(ints) - 1; low < high ; {
		for ;low < high; high-- {
			if ints[high] >= 0{
				break
			}
		}

		for ;low < high; low++ {
			if ints[low] <= 0{
				break
			}
		}

		// 处理元素都为0的情况
		if ints[low] == ints[high] {
			// 将low位置的元素与索引更大的非0元素对调
			for i := low ; i < high; i++{
				if ints[i] != 0{
					ints[low], ints[i] = ints[i], ints[low]
					break
				}
			}

			// 如果没有找到，说明low与high之间都为0
			if ints[low] == 0{
				return ints
			}

		}else {
			ints[high], ints[low] = ints[low], ints[high]
		}
	}

	return ints
}