package main

import "sort"

func longestPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	sort.Strings(s)
	first, last := s[0], s[len(s)-1]
	result := ""
	for i := 0; i < len(first); i++ {
		if i < len(last) && first[i] == last[i] {
			result += string(first[i])
		} else {
			break
		}
	}
	return result
}
