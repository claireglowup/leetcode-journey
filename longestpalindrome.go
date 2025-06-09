package main

func longestPalindrome(s string) string {
	check := func(left, right int) bool {
		left = left
		right = right - 1

		for left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}

		return true

	}

	for length := len(s); length > 0; length-- {
		for start := 0; start <= len(s)-length; start++ {
			if check(start, start+length) {
				return s[start : start+length]
			}
		}
	}

	return ""

}
