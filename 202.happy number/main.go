package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	seen := make(map[int]bool)
	for n != 1 && !seen[n] {
		seen[n] = true
		sum := 0
		for n > 0 {
			d := n % 10
			sum += d * d
			n /= 10
		}
		n = sum
	}
	return n == 1
}
