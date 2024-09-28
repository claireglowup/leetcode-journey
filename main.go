package main

import "fmt"

func main() {

	nums := []int{1, 11, 15, 7, 2}
	target := 9

	dict := make(map[int]int)

	for i, v := range nums {

		index, ok := dict[target-v]
		if ok {
			fmt.Println(index, i)
		}
		dict[v] = i
	}
}
