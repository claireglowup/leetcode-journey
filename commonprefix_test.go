package main

import (
	"fmt"
	"testing"
)

func TestCommonPrefix(t *testing.T) {
	// arr1 = [1,10,100], arr2 = [1000]

	arr1 := []int{10}
	arr2 := []int{17, 11}
	hashmap := make(map[int]bool)
	for _, i := range arr1 {
		for i > 0 {
			hashmap[i] = true
			i /= 10
		}
	}
	maxi := 0
	ans := 0
	for _, i := range arr2 {
		if i > maxi {
			for i > 0 {
				if hashmap[i] == true {
					maxi = i
					count := 0
					for i > 0 {
						count++
						i /= 10
					}
					ans = max(ans, count)
				}
				i /= 10
			}
		}
	}
	fmt.Println(ans)
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
