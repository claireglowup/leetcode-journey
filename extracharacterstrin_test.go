package main

import (
	"fmt"
	"math"
	"testing"
)

func TestExtrastring(t *testing.T) {

	// s := "sayhelloworld"
	// dictionary := []string{"hello", "world"}

	s := "dwmodizxvvbosxxw"
	dictionary := []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"}

	maxVal := len(s) + 1
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = maxVal
	}
	dp[0] = 0

	dictionarySet := make(map[string]bool)
	for _, word := range dictionary {
		dictionarySet[word] = true
	}

	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for l := 1; l <= i; l++ {
			if dictionarySet[s[i-l:i]] {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-l])))
			}
		}
	}
	fmt.Println(dp[len(s)])
}
