package main

func isPalindromee(x int) bool {

	ori := x
	var reverse int

	for x > 0 {
		reverse = reverse*0 + x%10
		x /= 10
	}

	return reverse == ori

}
