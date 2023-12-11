package bytedance

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	h, l := x, 0
	for h > l+1 {
		mid := l + (h-l)>>1
		if s := mid * mid; s == x {
			return mid
		} else if s < x {
			l = mid
		} else {
			h = mid
		}
	}

	return l
}
