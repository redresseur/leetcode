package number_operate

/*
	69 求数的平方根
	https://leetcode-cn.com/problems/sqrtx/
*/
import (
	"math"
)

// 牛顿求根公式
func mySqrt(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	r := n >> 1
	for r*r > n {
		r = (r + n/r) >> 1
	}

	return r
}

func mySqrt2(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	right, left, mid, mids := 0, n, 0, 0
	for right <= left {
		mid = (right + left) >> 1
		mids = mid * mid

		if mids == n {
			break
		} else if mids < n {
			if mids+(mid<<1)+1 > n {
				break
			}
			right = mid + 1
		} else {
			left = mid - 1
		}
	}

	return mid
}

func mySqrtFloat(n float64) float64 {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	fn := float64(n)
	r := fn / 2
	for i := 0; i < 1024; i++ {
		if r*r == fn {
			break
		}
		r = (r + fn/r) / 2
	}

	return r
}

//
//func mySqrtFloat2(n int) float64 {
//	if n < 0 {
//		return -1
//	} else if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	}
//
//	fn := float64(n)
//	right, left, mid, mids := float64(0), float64(n), float64(0), float64(0)
//
//	for offset := 0.0000001; right <= left; {
//		mid = (right + left) / 2
//		mids = mid * mid
//
//		if mids == fn {
//			break
//		} else if mids < fn {
//			if (mid+offset)*(mid+offset) > fn {
//				break
//			} else {
//				right = mid + offset
//			}
//		} else {
//
//			left = mid - offset
//
//		}
//	}
//
//	return mid
//}

//
//func mySqrtFloat2(n int) float64 {
//	if n < 0 {
//		return -1
//	} else if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	}
//
//	fn := float64(n)
//	right, left, mid, mids := float64(0), float64(n), float64(0), float64(0)
//
//	for i, offset := uint(0), float64(1); i <= 7 && right <= left; {
//		mid = (right + left) / 2
//		mids = fn / mid
//
//		if mids == mid {
//			break
//		} else if mids > mid {
//			if mid + offset > mids {
//				if offset == 0.0000001 {
//					break
//				} else {
//					offset = 0.0000001
//				}
//			}
//
//			right = mid + offset
//
//		} else {
//			if mid - offset < mids {
//				if offset == 0.0000001 {
//					break
//				} else {
//					offset = 0.0000001
//				}
//			}
//			left = mid - offset
//		}
//	}
//
//	return mid
//}

func mySqrtFloat3(n float64) float64 {
	right, left, mid, mids := float64(1), float64(n), float64(0), float64(0)
	if n < 1 {
		right = n
		left = 1
	}

	mid = right + (left-right)/2
	for {
		mids = mid * mid
		if math.Abs(mids-n) > 0.000001 {
			if mids < n {
				right = mid
			} else {
				left = mid
			}

			mid = (right + left) / 2
		} else {
			break
		}
	}

	return mid
}

func mySqrtFloat2(n float64) float64 {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	// 牛顿迭代公式
	r := int64(n) >> 1
	if r > 0 {
		for {
			r = (r + int64(n)/r) >> 1
			if rr := r * r; rr == int64(n) {
				return float64(r)
			} else if rr < int64(n) {
				break
			}
		}
	}

	//right, left, mid, mids := float64(r), float64(r+1), float64(0), float64(0)
	right, left, mid, mids := float64(0), float64(n), float64(0), float64(0)

	for i, offset := uint(0), math.Exp2(-32); right <= left; i++ {
		//fmt.Println(i)
		mid = (right + left) / 2
		mids = mid * mid

		if mids == n {
			break
		} else if mids < n {
			if (mid+offset)*(mid+offset) > n {
				//right = mid
				//				//left = mid + offset
				//				//i++
				//				//offset /= 10
				break
			} else {
				right = mid + offset
			}
		} else {
			//if (mid-offset)*(mid-offset) < n {
			//	right = mid - offset
			//	left = mid
			//	i++
			//	offset /= 10
			//} else {
			left = mid - offset
			//}
		}
	}

	return mid
}

//
//func mySqrtFloat2(n float64) float64 {
//	if n < 0 {
//		return -1
//	} else if n == 0 {
//		return 0
//	} else if n == 1 {
//		return 1
//	}
//
//	// 牛顿迭代公式
//	r := int64(n) >> 1
//	if r > 0 {
//		for {
//			r = (r + int64(n)/r) >> 1
//			if rr := r * r; rr == int64(n) {
//				return float64(r)
//			} else if rr < int64(n) {
//				break
//			}
//		}
//	}
//
//	right, left, mid, mids := float64(r), float64(r+1), float64(0), float64(0)
//
//	for i, offset := uint(0), float64(0.1); i <= 7 && right <= left; {
//		mid = (right + left) / 2
//		mids = mid * mid
//
//		if mids == n {
//			break
//		} else if mids < n {
//			if (mid+offset)*(mid+offset) > n {
//				right = mid
//				left = mid + offset
//				i++
//				offset /= 10
//			} else {
//				right = mid + offset
//			}
//		} else {
//			if (mid-offset)*(mid-offset) < n {
//				right = mid - offset
//				left = mid
//				i++
//				offset /= 10
//			} else {
//				left = mid - offset
//			}
//		}
//	}
//
//	return mid
//}
