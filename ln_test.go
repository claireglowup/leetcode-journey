package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// nums =
// [10,2,9,39,17]

// Use Testcase
// Output
// "92103917"
// Expected

func TestLargestNumber(t *testing.T) {

	// nums := []int{3, 30, 34, 5, 9} "9534330"
	// nums := []int{2, 10} 210
	nums := []int{10, 2, 9, 39, 17} // "93921710"

	vs := make([]string, len(nums))
	for i, v := range nums {
		vs[i] = strconv.Itoa(v)
	}

	sort.Slice(vs, func(i, j int) bool {
		return vs[i]+vs[j] > vs[j]+vs[i]
	})

	if vs[0] == "0" {
		return
	}

	result := strings.Join(vs, "")
	fmt.Println(result)

}
