package longestsubstirng

import "strings"

// 编号： 3
//Given a string, find the length of the longest substring without repeating characters.
//
//Example 1:
//
//Input: "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
func lengthOfLongestSubstring(str string) (res int) {
	sub := ""
	for i := 0; i < len(str); i++ {
		c := string(str[i])
		if index := strings.Index(sub, c); index < 0 {
			sub += c
		} else {
			if res < len(sub) {
				res = len(sub)
			}

			sub = sub[index+1:] + c
		}
	}

	if res < len(sub) {
		res = len(sub)
	}

	return
}
