package main

import "log"

func main() {

	log.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {

	ori := x
	var reverse int

	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10

	}
	return reverse == ori

}
