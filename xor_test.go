package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorQueriesSubArray(t *testing.T) {

	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	var result []int
	var xori int
	for _, query := range queries {
		xori = arr[query[0]]
		for i := query[0] + 1; i <= query[1]; i++ {
			xori ^= arr[i]
		}
		result = append(result, xori)
	}

	assert.Equal(t, []int{2, 7, 14, 8}, result)
}
