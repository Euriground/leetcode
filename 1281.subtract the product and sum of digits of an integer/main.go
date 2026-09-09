package main

import "fmt"

func main() {
	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for n > 0 {
		d := n % 10
		product *= d
		sum += d
		n /= 10
	}
	return product - sum
}
