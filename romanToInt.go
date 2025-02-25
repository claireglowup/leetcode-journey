package main

func romanToInt(s string) int {

	roman := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	var total, prev, current int

	for _, char := range s {
		current = roman[char]
		if prev < current {
			total += current - 2*prev
		} else {
			total += current
		}

		prev = current
	}

	return total
}
