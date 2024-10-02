package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5

	var result [][]int
	remainderMap := make(map[int]int) // Map untuk menyimpan sisa pembagian

	for _, each := range arr {
		rem := each % k
		complement := (k - rem) % k // Cari pasangan yang melengkapi jadi kelipatan k

		if remainderMap[complement] > 0 {
			result = append(result, []int{each, complement * k})
			remainderMap[complement]-- // Kurangi jumlah pasangan yang tersedia
		} else {
			remainderMap[rem]++ // Tambahkan elemen ke map
		}
	}

	fmt.Println(result)
}
