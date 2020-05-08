package longestPdrmcSubstr

// 编号： 5
//Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
//
//Example 1:
//
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.

func getRange(s string, i, j int, up, down *int) {
	ti, tj := 0, 0
	for ; i >= 0 && j <= len(s)-1; i, j = i-1, j+1 {
		if s[i] != s[j] {
			break
		}

		ti = i
		tj = j
	}

	if tj-ti > *up-*down {
		*up = tj
		*down = ti
	}
}

func longestPalindrome(s string) (res string) {
	up, down, lenStr := 0, 0, len(s)
	if 0 == lenStr {
		return
	}

	for i := 0; i < lenStr-1; i++ {
		if i < lenStr-2 {
			getRange(s, i, i+2, &up, &down)
		}

		if i < lenStr-1 {
			getRange(s, i, i+1, &up, &down)
		}

		if up >= lenStr-1 {
			break
		}
	}

	res = s[down : up+1]
	return
}
