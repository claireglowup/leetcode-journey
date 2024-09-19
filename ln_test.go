package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLargestNumber(t *testing.T) {

	// nums := []int{3, 30, 34, 5, 9} "9534330"
	nums := []int{2, 10} //210

	for i := 1; i < len(nums); i++ {
		cek, _ := strconv.Atoi(strconv.Itoa(nums[i-1]) + strconv.Itoa(nums[i]))
		cek2, _ := strconv.Atoi(strconv.Itoa(nums[i]) + strconv.Itoa(nums[i-1]))

		if cek < cek2 {
			nums[i] = cek2
		} else {
			nums[i] = cek

		}

	}
	fmt.Println(nums[len(nums)-1])

}
